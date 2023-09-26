package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type CasbinRouter struct{}

func (cr *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	casbinRouter := Router.Group("casbin")
	casbinApi := api.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.PUT("", casbinApi.EditCasbin)
	}
	return casbinRouter
}
