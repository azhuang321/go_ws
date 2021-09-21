package main

import (
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"user_srv/handler"
	"user_srv/initialize"
	"user_srv/proto/gen/userpb"
)

func initFrameWork() {
	initialize.InitConfig()
	initialize.InitLogger()
	initialize.InitDB()
}

func main() {
	initFrameWork()

	gs := grpc.NewServer()

	userService := handler.UserService{}
	userpb.RegisterUserServer(gs, userService)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		zap.S().Errorf("监听端口失败:%s\n", err.Error())
		return
	}
	waitChan := make(chan interface{})

	go func() {
		err = gs.Serve(lis)
		if err != nil {
			waitChan <- err
			zap.S().Errorf("服务启动失败:%s\n", err.Error())
			return
		}
	}()
	zap.S().Infof("服务启动成功")

	<-waitChan

}
