package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type CasbinRouter struct{}

func (cr *CasbinRouter) InitCasbinRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	casbinRouter := Router.Group("casbin").Use(middleware.OperationRecord())
	casbinApi := api.ApiGroupApp.SystemApiGroup.CasbinApi
	{
		casbinRouter.PUT("", casbinApi.EditCasbin)
	}
	return casbinRouter
}
