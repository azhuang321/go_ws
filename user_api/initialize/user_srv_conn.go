package initialize

import (
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"user_api/global"
	"user_api/proto/gen/go/userpb"
)

func InitUserSrvConn() {
	consulInfo := global.Config.Consul

	// 通过负载均衡器 去注册中心拿用户服务
	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s", consulInfo.Host, consulInfo.Port, global.Config.UserSrvName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[InitUserSrvConn] 连接 [user_srv] 失败:%s",  err.Error())
		return
	}
	userSrvClient := userpb.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}