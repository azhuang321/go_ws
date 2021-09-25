package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"ws_srv/config"
	"ws_srv/proto/gen/go/msgpb"
	"ws_srv/proto/gen/go/userpb"
)

var (
	DB *gorm.DB

	GClient *redis.Client

	Config  *config.Config

	UserSrvClient userpb.UserClient

	SocketRouter = map[string]func(*websocket.Conn,*msgpb.Msg)error{}
)
