package global

import (
	"chat_api/config"
	"chat_api/proto/gen/go/chat_pb"
	"chat_api/proto/gen/go/userpb"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB

	GClient *redis.Client

	Config *config.Config

	ChatSrvClient chat_pb.ChatClient

	UserSrvClient userpb.UserClient
)
