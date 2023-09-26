package system

import (
	"server/global"
)

type OperationRecord struct {
	global.MODEL
	Menu     string `json:"menu" gorm:"not null;column:menu;comment:操作菜单"`             // 操作菜单
	Method   string `json:"method" gorm:"not null;column:method;comment:请求方法"`         // 请求方法
	Uri      string `json:"uri" gorm:"column:path;comment:uri"`                        // uri
	Status   int    `json:"status" gorm:"not null;column:status;comment:请求状态"`         // 状态码
	ReqParam string `json:"reqParam" gorm:"type:text;column:req_param;comment:请求Body"` // 请求参数
	RespData string `json:"respData" gorm:"type:text;column:resp_data;comment:响应数据"`   // 响应数据
	RespTime int64  `json:"responseTime" gorm:"column:responseTime"`                   // 响应时长
	UserName string `json:"username" gorm:"not null;user_name;comment:用户名称"`           // 用户名称
}

func (o *OperationRecord) TableName() string {
	return "sys_operation_record"
}
