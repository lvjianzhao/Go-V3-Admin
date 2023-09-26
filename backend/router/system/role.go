package system

import (
	"github.com/gin-gonic/gin"
	"server/api"
	"server/middleware"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	roleRouter := Router.Group("role").Use(middleware.OperationRecord("角色管理"))
	roleApi := api.ApiGroupApp.SystemApiGroup.RoleApi
	{
		roleRouter.GET("", roleApi.GetRoles)
		roleRouter.POST("", roleApi.AddRole)
		roleRouter.DELETE("/:roleID", roleApi.DeleteRole)
		roleRouter.PUT("", roleApi.EditRole)
		roleRouter.PUT("roleMenu", roleApi.EditRoleMenu)
	}
	return roleRouter
}
