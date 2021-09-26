package handler

import (
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
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
	toUserId := msg.GetContent().GetTo().Id
	fmt.Println(toUserId)
	toUserConn := api.UserClientConn[toUserId]
	fmt.Println(toUserConn)

	err := toUserConn.UserConn.WriteMessage(websocket.BinaryMessage,[]byte("123"))
	if err != nil {
		zap.S().Errorf("发送消息失败:%s",err.Error())
	}
	fmt.Println("send success")
	return nil
}
