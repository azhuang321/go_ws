package initialize

import (
	"ws_srv/global"
	"ws_srv/handler"
)

func InitSocketRouter() {
	global.SocketRouter["/chat"] = handler.Chat
}
