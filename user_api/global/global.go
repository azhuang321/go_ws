package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"user_api/config"
)

var (
	DB *gorm.DB

	GClient *redis.Client

	Config = &config.Config{}
)
