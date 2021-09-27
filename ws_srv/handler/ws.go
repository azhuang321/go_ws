package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"ws_srv/api"
	"ws_srv/proto/gen/go/msgpb"
	"ws_srv/utils"
)

func Auth(conn *websocket.Conn,msg *msgpb.Msg) error {
	jwt := utils.NewJWT()
	claims,err := jwt.ParseToken(msg.GetToken())
	if err != nil {
		zap.S().Errorf("解析token失败:%s",err.Error())
		conn.Close()
		return err
	}
	zap.S().Infof("登录会员信息:%#v",claims)
	api.UserClientConn[claims.ID] = &api.UserConn{
		UserClaims: claims,
		UserConn:   conn,
	}

	if err != nil {
		zap.S().Errorf("发送消息失败:%s",err.Error())
	}
	return nil
}

func Chat(conn *websocket.Conn,msg *msgpb.Msg) error {
	fmt.Printf("%#v\n",msg.GetContent().GetSendInfo().GetSendUserInfo())
	toUserId := msg.GetContent().GetReceiveInfo().GetReceiveUserInfo().GetId()
	toUserConn := api.UserClientConn[toUserId]

	binary,err := proto.Marshal(msg)
	if err != nil {
		zap.S().Errorf("序列化失败:%s",err.Error())
		return err
	}

	//todo 判断,修改为sync,map

	err = toUserConn.UserConn.WriteMessage(websocket.BinaryMessage,binary)
	if err != nil {
		zap.S().Errorf("发送消息失败:%s",err.Error())
	}
	fmt.Println("send success")
	return nil
}
