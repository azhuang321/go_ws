package initialize

import (
	"chat_api/proto/gen/go/chat_pb"
	"fmt"

	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"chat_api/global"
)

func InitChatSrvConn() {
	consulInfo := global.Config.Consul

	// 通过负载均衡器 去注册中心拿用户服务
	chatConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s", consulInfo.Host, consulInfo.Port, global.Config.ChatSrvName),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	)
	if err != nil {
		zap.S().Errorf("[InitUserSrvConn] 连接 [user_srv] 失败:%s", err.Error())
		return
	}
	userSrvClient := chat_pb.NewChatClient(chatConn)
	global.ChatSrvClient = userSrvClient
}
