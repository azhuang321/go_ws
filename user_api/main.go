package main

import (
	"user_api/core"
	"user_api/initialize"
)

func initFrameWork() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
	initialize.InitTranslate()
	initialize.InitRedisPool()
}

func main() {
	initFrameWork()
	//gorm.io/gen
	core.StartServer()
}
