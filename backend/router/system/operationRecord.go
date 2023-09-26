package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type OperationRecordRouter struct{}

func (o *OperationRecordRouter) InitOperationRecordRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	operationRecordRouter := Router.Group("record")
	operationRecordApi := api.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.GET("", operationRecordApi.GetOperationRecordList)
		operationRecordRouter.GET("/options", operationRecordApi.GetSearchOptions)
	}
	return operationRecordRouter
}
