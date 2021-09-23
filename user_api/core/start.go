package core

import (
	"fmt"

	"go.uber.org/zap"

	"user_api/initialize"
)

func StartServer() {
	port := 9001
	waitChan := make(chan interface{})
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