package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user").Use(middleware.OperationRecord())
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.GET("detail", userApi.GetUserInfo)
		userRouter.GET("", userApi.GetUsers)
		userRouter.DELETE("/:userID", userApi.DeleteUser)
		userRouter.POST("", userApi.AddUser)
		userRouter.PUT("", userApi.EditUser)
		userRouter.PUT("password", userApi.ModifyPass)
		userRouter.PUT("status", userApi.SwitchActive)
	}
	return userRouter
}
