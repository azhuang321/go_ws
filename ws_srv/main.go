package main

import (
	"fmt"

	"go.uber.org/zap"

	"ws_srv/initialize"
)

func initFrameWork() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

func main() {
	initFrameWork()
	port := 9001
	waitChan := make(chan interface{})
	//gorm.io/gen

	// 初始化router
	Router := initialize.Routers()
	go func() {
		zap.S().Infof("启动web服务，端口：%d", port)
		if err := Router.Run(fmt.Sprintf(":%d", port)); err != nil {
			waitChan <- err
			zap.S().Panic("启动web服务失败：", err.Error())
		}
	}()

	<-waitChan
}
