package system

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"server/utils"

	"server/global"
	"server/model/common/response"
	systemReq "server/model/system/request"
)

type OperationRecordApi struct{}

// GetOperationRecordList 分页获取操作记录
func (o *OperationRecordApi) GetOperationRecordList(c *gin.Context) {

	var orSp systemReq.OrSearchParams
	orSp.Page, _ = utils.StringToInt(c.Query("currentPage"))
	orSp.PageSize, _ = utils.StringToInt(c.Query("size"))
	orSp.Status = c.Query("status")
	orSp.Method = c.Query("method")
	orSp.Username = c.Query("username")
	orSp.Menu = c.Query("menu")
	orSp.StartTime, _ = utils.StringToInt64(c.Query("startTime"))
	orSp.EndTime, _ = utils.StringToInt64(c.Query("endTime"))

	if list, total, err := operationService.GetOperationRecordList(orSp); err != nil {
		response.FailWithMessage("获取失败", c)
		global.LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     orSp.Page,
			PageSize: orSp.PageSize,
		}, "获取成功", c)
	}
}

// GetSearchOptions 获取搜索条件的选项
func (o *OperationRecordApi) GetSearchOptions(c *gin.Context) {
	if options, err := operationService.GetSearchOptions(); err != nil {
		response.FailWithMessage("获取失败", c)
		global.LOG.Error("获取失败", zap.Error(err))
	} else {
		response.OkWithDetailed(response.PageResult{
			List: options,
		}, "获取成功", c)
	}
}
