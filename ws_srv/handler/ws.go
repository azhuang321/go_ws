package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"ws_srv/logic"
	"ws_srv/proto/gen/go/msgpb"
)

func Chat(conn *websocket.Conn,msg *msgpb.Msg) error {
	toUserId := msg.GetContent().GetReceiveInfo().GetReceiveUserInfo().GetId()
	zap.S().Infof("接受者id:%d",toUserId)

	toUserConn,ok := logic.UserClientConn.Load(toUserId)
	if !ok {
		zap.S().Errorf("收消息者不在线: %d", toUserId)
		return nil
	}

	fmt.Println("消息",msg.GetContent().GetReceiveInfo().GetContent())

	binary,err := proto.Marshal(msg)
	if err != nil {
		zap.S().Errorf("序列化失败:%s",err.Error())
		return err
	}

	err = toUserConn.(*websocket.Conn).WriteMessage(websocket.BinaryMessage,binary)
	if err != nil {
		zap.S().Errorf("发送消息失败:%s",err.Error())
		return err
	}
	fmt.Println("send success")
	return nil
}
