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
		apiRouter.POST("addApi", apiApi.AddApi)
		apiRouter.GET("getApis", apiApi.GetApis)
		apiRouter.GET("apiGroups", apiApi.GetApiGroups)
		apiRouter.POST("deleteApi", apiApi.DeleteApi)
		apiRouter.POST("editApi", apiApi.EditApi)
		apiRouter.POST("getElTreeApis", apiApi.GetElTreeApis)
	}
	return apiRouter
}
