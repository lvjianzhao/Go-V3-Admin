package system

import "server/global"

type ApiModel struct {
	global.TD27_MODEL
	Path        string `json:"path" gorm:"comment:api路径" validate:"required"`             // api路径
	Description string `json:"description" gorm:"comment:api中文描述" validate:"required"`    // api中文描述
	Group       string `json:"group" gorm:"comment:api组" validate:"required"`             // api组
	Method      string `json:"method" gorm:"default:POST;comment:方法" validate:"required"` // 方法:创建POST(默认)|查看GET|更新PUT|删除DELETE
}

func (ApiModel) TableName() string {
	return "sys_api"
}
