package main

import (
	"chat_api/core"
	"chat_api/initialize"
)

func initFrameWork() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitTranslate()
	initialize.InitRedisPool()
	initialize.InitChatSrvConn()
	initialize.InitUserSrvConn()
}

func main() {
	initFrameWork()

	core.StartServer()
}
