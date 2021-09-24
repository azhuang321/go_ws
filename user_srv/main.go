package main

import (
	"user_srv/core"

	"user_srv/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()

	core.RegisterService()
	core.MainExit()

}
