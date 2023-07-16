package internal

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path"
	"server/global"
)

type lumberjackLogs struct{}

var LumberjackLogs = new(lumberjackLogs)

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (l *lumberjackLogs) GetWriteSyncer(level string) zapcore.WriteSyncer {
	fileWriter := &lumberjack.Logger{
		Filename:   path.Join(global.CONFIG.Zap.Director, "server.log"),
		MaxSize:    global.CONFIG.RotateLogs.MaxSize,
		MaxBackups: global.CONFIG.RotateLogs.MaxBackups,
		MaxAge:     global.CONFIG.RotateLogs.MaxAge,
		Compress:   global.CONFIG.RotateLogs.Compress,
	}

	if global.CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter))
	}
	return zapcore.AddSync(fileWriter)
}
