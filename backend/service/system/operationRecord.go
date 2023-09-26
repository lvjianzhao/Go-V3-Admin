package system

import (
	"fmt"
	"server/global"
	modelSystem "server/model/system"
	systemReq "server/model/system/request"
	"strings"
	"time"
)

type OperationRecordService struct{}

type Options struct {
	Value string `json:"value"`
	Label string `json:"label"`
}

type SearchOptions struct {
	Menu     []*Options `json:"menu"`
	UserName []*Options `json:"user_name"`
	Status   []*Options `json:"status"`
}

// CreateOperationRecord 创建记录
func (o *OperationRecordService) CreateOperationRecord(operationRecord modelSystem.OperationRecord) error {
	return global.DB.Create(&operationRecord).Error
}

// GetOperationRecordList 分页获取操作记录
func (o *OperationRecordService) GetOperationRecordList(orSp systemReq.OrSearchParams) ([]modelSystem.OperationRecord, int64, error) {
	var orList []modelSystem.OperationRecord
	var total int64

	limit := orSp.PageSize
	offset := orSp.PageSize * (orSp.Page - 1)
	db := global.DB.Model(&modelSystem.OperationRecord{}).
		Order("created_at DESC")

	if orSp.Method != "" {
		methodSlice := strings.Split(orSp.Method, ",")
		methodOrConditions := make([]string, len(methodSlice))
		args := make([]interface{}, len(methodSlice))
		for i, method := range methodSlice {
			// 在这里处理每个 method 的逻辑
			methodOrConditions[i] = "method = ?"
			args[i] = method
		}
		db = db.Where(strings.Join(methodOrConditions, " OR "), args...)
	}

	if orSp.Menu != "" {
		menuSlice := strings.Split(orSp.Menu, ",")
		menuOrConditions := make([]string, len(menuSlice))
		args := make([]interface{}, len(menuSlice))
		for i, menu := range menuSlice {
			// 在这里处理每个 menu 的逻辑
			menuOrConditions[i] = "menu = ?"
			args[i] = menu
		}
		db = db.Where(strings.Join(menuOrConditions, " OR "), args...)
	}

	if orSp.Status != "" {
		statusSlice := strings.Split(orSp.Status, ",")
		statusOrConditions := make([]string, len(statusSlice))
		args := make([]interface{}, len(statusSlice))
		for i, status := range statusSlice {
			// 在这里处理每个 status 的逻辑
			statusOrConditions[i] = "status = ?"
			args[i] = status
		}
		db = db.Where(strings.Join(statusOrConditions, " OR "), args...)
	}

	if orSp.Username != "" {
		db.Where("user_name = ?", orSp.Username)
	}

	if orSp.StartTime != 0 && orSp.EndTime != 0 {
		startTime := time.Unix(orSp.StartTime, 0)
		endTime := time.Unix(orSp.EndTime, 0)
		db.Where("created_at BETWEEN ? AND ?", startTime, endTime)
	}

	err := db.Count(&total).Error

	if err != nil {
		return orList, total, fmt.Errorf("查询 OperationRecord 失败: %v", err)
	}

	db.Limit(limit).Offset(offset).Find(&orList)
	return orList, total, err
}

// GetSearchOptions 获取搜索表单需要的选项列表
func (o *OperationRecordService) GetSearchOptions() (so SearchOptions, err error) {
	var menuOptions []string
	var usernameOptions []string
	var statusOptions []string

	err = global.DB.
		Model(&modelSystem.OperationRecord{}).
		Distinct("menu").
		Pluck("menu", &menuOptions).
		Error

	if err != nil {
		return
	}

	err = global.DB.
		Model(&modelSystem.OperationRecord{}).
		Distinct("user_name").
		Pluck("user_name", &usernameOptions).
		Error

	if err != nil {
		return
	}

	err = global.DB.
		Model(&modelSystem.OperationRecord{}).
		Distinct("status").
		Pluck("status", &statusOptions).
		Error

	if err != nil {
		return
	}

	// 将查询结果中的值添加到切片中
	for _, menu := range menuOptions {
		so.Menu = append(so.Menu, &Options{
			Value: menu,
			Label: menu,
		})
	}

	for _, user := range usernameOptions {
		so.UserName = append(so.UserName, &Options{
			Value: user,
			Label: user,
		})
	}

	for _, s := range statusOptions {
		so.Status = append(so.Status, &Options{
			Value: s,
			Label: s,
		})
	}
	return
}
