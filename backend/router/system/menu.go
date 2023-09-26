package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type MenuRouter struct{}

func (u *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	menuRouter := Router.Group("menu").Use(middleware.OperationRecord("菜单管理"))
	menuApi := api.ApiGroupApp.SystemApiGroup.MenuApi
	{
		menuRouter.GET("", menuApi.GetMenus)
		menuRouter.POST("", menuApi.AddMenu)
		menuRouter.PUT("", menuApi.EditMenu)
		menuRouter.DELETE("/:menuID", menuApi.DeleteMenu)
		menuRouter.GET("getElTreeMenus", menuApi.GetElTreeMenus)
	}
	return menuRouter
}
