package main

import (
	"ws_srv/core"
	"ws_srv/initialize"
)

func initFrameWork() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitTranslate()
	initialize.InitRedisPool()
	initialize.InitUserSrvConn()
	initialize.InitSocketRouter()
}

func main() {
	initFrameWork()

	core.StartServer()
}
