package handler

import (
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"ws_srv/api"
	"ws_srv/utils"
	"ws_srv/proto/gen/go/msgpb"
)



func Auth(conn *websocket.Conn,msg *msgpb.Msg) error {
	jwt := utils.NewJWT()
	//claims := &utils.CustomClaims{}
	claims,err := jwt.ParseToken(msg.GetToken())
	if err != nil {
		zap.S().Errorf("解析token失败:%s",err.Error())
		conn.Close()
		return err
	}
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


	return nil

}
