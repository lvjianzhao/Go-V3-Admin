package api

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"server/core"
	"server/global"
	"server/initialize"
	"server/pkg/crontab"
	"sync"
)

var (
	wg       sync.WaitGroup
	config   string
	StartCmd = &cobra.Command{
		Use:     "serve",
		Short:   "Start API server",
		Example: "go-admin serve",
		PreRun: func(cmd *cobra.Command, args []string) {
			usage()
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func init() {
	StartCmd.PersistentFlags().StringVarP(&config, "config", "c", "config.yaml", "Start server with provided configuration file")
}

func usage() {
	usageStr := `starting api server`
	fmt.Printf("%s\n", usageStr)
}

func setup() {
	// 1. 读取配置
	global.TD27_VP = core.Viper(config) // 初始化viper
	// 2. 初始化日志以及数据库连接
	global.TD27_LOG = core.Zap() // 初始化zap日志
	zap.ReplaceGlobals(global.TD27_LOG)
	global.TD27_DB = initialize.Gorm() // gorm连接数据库

	if global.TD27_DB == nil {
		global.TD27_LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.TD27_DB) // 初始化表
	}

	// 3. 开启定时任务
	global.TD27_CRON = cron.New()
	crontab.StartCrontab()
}

func run() error {
	// 程序结束前关闭数据库连接以及停止定时任务
	db, _ := global.TD27_DB.DB()
	defer db.Close()
	defer global.TD27_CRON.Stop()

	core.RunServer()
	return nil
}
