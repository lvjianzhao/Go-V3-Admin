package crontab

import (
	"fmt"
	"reflect"
	"runtime"
	"server/global"
)

// checkAndUpdateServiceStatus 更新维保状态以及检查一个月内即将到期的维保客户
func checkAndUpdateServiceStatus() {
	global.TD27_LOG.Info("定时任务 【xxx】 开始执行...")
	// 执行定时任务的代码
	global.TD27_LOG.Info("定时任务 【xxx】 执行完成...")
}

// addCron 添加一个定时任务
func addCron(spec string, cmd func()) {
	pc := reflect.ValueOf(cmd).Pointer()
	funcName := runtime.FuncForPC(pc).Name()

	id, err := global.TD27_CRON.AddFunc(spec, cmd)
	if err != nil {
		global.TD27_LOG.Info(
			fmt.Sprintf("Add Cron Failure,spec: %s,cron: %s,error: %s", spec, funcName, err))
	} else {
		global.TD27_LOG.Info(fmt.Sprintf("Add Crontab Success; ID: %v,task: %s", id, funcName))
	}
}

func StartCrontab() {
	addCron("* 7 * * *", checkAndUpdateServiceStatus)
	global.TD27_CRON.Start()
}
