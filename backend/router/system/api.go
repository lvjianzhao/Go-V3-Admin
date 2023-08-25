package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type ApiRouter struct{}

func (u *ApiRouter) InitApiRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiApi := api.ApiGroupApp.SystemApiGroup.ApiApi
	{
		apiRouter.POST("", apiApi.AddApi)
		apiRouter.GET("", apiApi.GetApis)
		apiRouter.DELETE("/:apiID", apiApi.DeleteApi)
		apiRouter.PUT("", apiApi.EditApi)
		apiRouter.GET("apiGroups", apiApi.GetApiGroups)
		apiRouter.GET("getElTreeApis", apiApi.GetElTreeApis)
		// 测试接口
		apiRouter.GET("table", apiApi.GetTable)
	}
	return apiRouter
}
