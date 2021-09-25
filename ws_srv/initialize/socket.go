package initialize

import (
	"ws_srv/global"
	"ws_srv/handler"
)

func InitSocketRouter() {
	global.SocketRouter["/auth"] = handler.Auth
	global.SocketRouter["/chat"] = handler.Chat
}
