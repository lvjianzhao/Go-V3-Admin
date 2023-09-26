package core

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"

	"server/global"
	"server/initialize"
)

func RunServer() {
	if global.CONFIG.System.UseMultipoint {
		initialize.Redis()
	}

	addr := fmt.Sprintf("%s:%d", global.CONFIG.System.Host, global.CONFIG.System.Port)
	router := initialize.Routers()
	srv := http.Server{
		Addr:           addr,
		Handler:        router,
		ReadTimeout:    120 * time.Second,
		WriteTimeout:   120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.LOG.Error("listen", zap.Error(err))
			srv.Shutdown(ctx)
			os.Exit(1)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.LOG.Info("Shutdown Server ...")

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Error("Server Shutdown", zap.Error(err))
	}
	global.LOG.Info("Server exiting")
}
