package request

import "server/model/common/request"

// OrSearchParams 操作记录搜索条件
type OrSearchParams struct {
	Method    string `json:"method"`
	Menu      string `json:"menu"`
	Username  string `json:"username"`
	Status    string `json:"status"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
	request.PageInfo
}
