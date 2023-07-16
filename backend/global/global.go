package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"

	"server/config"
)

var (
	VP                  *viper.Viper
	CONFIG              config.Server
	LOG                 *zap.Logger
	DB                  *gorm.DB
	CRON                *cron.Cron
	REDIS               *redis.Client
	Concurrency_Control = &singleflight.Group{}
)
