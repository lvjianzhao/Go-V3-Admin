package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/resources"
)

type WebHandler struct{}

// RedirectIndex 重定向
func (w *WebHandler) RedirectIndex(c *gin.Context) {
	c.Redirect(http.StatusFound, "/ui/")
	return
}

func (w *WebHandler) Index(c *gin.Context) {
	c.Header("content-type", "text/html;charset=utf-8")
	c.String(200, string(resources.Html))
	return
}
