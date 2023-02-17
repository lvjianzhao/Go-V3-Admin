package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi
	{
		userRouter.POST("getUserInfo", userApi.GetUserInfo)
	}
	return userRouter
}