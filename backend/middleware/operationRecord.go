package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"

	"server/global"
	systemModel "server/model/system"
	"server/service"
	"server/utils"
)

var (
	operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService
)

type responseProxyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseProxyWriter) Write(b []byte) (int, error) {
	if r.body == nil {
		r.body = bytes.NewBufferString("")
	}
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func OperationRecord(menu string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求参数
		var reqParam []byte
		// get 请求不记录
		if c.Request.Method == http.MethodGet {
			return
			//query := c.Request.URL.RawQuery
			//query, _ = url.QueryUnescape(query)
			//split := strings.Split(query, "&")
			//m := make(map[string]string)
			//for _, v := range split {
			//	kv := strings.Split(v, "=")
			//	if len(kv) == 2 {
			//		m[kv[0]] = kv[1]
			//	}
			//}
			//reqParam, _ = json.Marshal(&m)
		} else {
			var err error
			reqParam, err = io.ReadAll(c.Request.Body)
			if err != nil {
				global.LOG.Error("read body from request error:", zap.Error(err))
			} else {
				c.Request.Body = io.NopCloser(bytes.NewBuffer(reqParam))
			}
		}

		// 解析token
		claims, _ := utils.GetClaims(c)

		record := systemModel.OperationRecord{
			Menu:     menu,
			Method:   c.Request.Method,
			Uri:      c.Request.URL.Path,
			ReqParam: string(reqParam),
			UserName: claims.Username,
		}

		writer := responseProxyWriter{
			ResponseWriter: c.Writer,
			body:           &bytes.Buffer{},
		}
		c.Writer = writer
		now := time.Now()

		c.Next()

		record.Status = c.Writer.Status()
		record.RespTime = time.Since(now).Milliseconds()
		record.RespData = writer.body.String()

		if err := operationRecordService.CreateOperationRecord(record); err != nil {
			global.LOG.Error("create operation record error:", zap.Error(err))
		}
	}
}
