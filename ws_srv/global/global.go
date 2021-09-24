package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"ws_srv/config"
	"ws_srv/proto/gen/go/userpb"
)

var (
	DB *gorm.DB

	GClient *redis.Client

	Config  *config.Config

	UserSrvClient userpb.UserClient
)
