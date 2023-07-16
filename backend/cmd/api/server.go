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
)

var (
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
	global.VP = core.Viper(config) // 初始化viper
	// 2. 初始化日志以及数据库连接
	global.LOG = core.Zap() // 初始化zap日志
	zap.ReplaceGlobals(global.LOG)
	global.DB = initialize.Gorm() // gorm连接数据库

	if global.DB == nil {
		global.LOG.Error("mysql连接失败，退出程序")
		os.Exit(127)
	} else {
		initialize.RegisterTables(global.DB) // 初始化表
	}

	// 3. 开启定时任务
	global.CRON = cron.New()
	crontab.StartCrontab()
}

func run() error {
	// 程序结束前关闭数据库连接以及停止定时任务
	db, _ := global.DB.DB()
	defer db.Close()
	defer global.CRON.Stop()

	core.RunServer()
	return nil
}
