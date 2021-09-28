package main

import (
	"chat_srv/core"

	"chat_srv/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	core.RegisterService()
	core.MainExit()
}
