package logic

import (
	"errors"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"net/http"
	"sync"
	"ws_srv/proto/gen/go/msgpb"
)
var UpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var UserClientConn = &sync.Map{}

type Message struct{
	UserId uint32
	Conn *websocket.Conn
	Msg *msgpb.Msg
}

// OnConnect 连接上系统需要处理相关
func OnConnect() {

}

// CheckAndUnmarshalMsg 检查验证消息
func CheckAndUnmarshalMsg(messageType int,receive []byte) (*msgpb.Msg,error) {
	if messageType != websocket.BinaryMessage {
		return nil,errors.New("消息类型不正确")
	}
	msg := &msgpb.Msg{}
	err := proto.Unmarshal(receive,msg)
	if err != nil {
		zap.S().Errorf("解析消息失败:%s",err.Error())
		return nil,errors.New("解析消息失败")
	}
	return msg,nil
}

// ConnClose 断开连接处理
func ConnClose(conn *websocket.Conn,userId uint32){
	zap.S().Info("断线处理")
	err := conn.Close()
	if err != nil {
		zap.S().Errorf("断开连接处理失败:%s",err.Error())
	}
	UserClientConn.Delete(userId)
}

// CloseConn 主动关闭
func CloseConn(userId uint32,reason string){
	userConn,ok := UserClientConn.Load(userId)
	if !ok {
		zap.S().Errorf("当关闭客户端时,未找到需要关闭的conn,userId:%d:",userId)
		return
	}
	conn := userConn.(*websocket.Conn)
	err := conn.WriteMessage(websocket.CloseMessage,websocket.FormatCloseMessage(websocket.CloseNormalClosure, reason))
	if err != nil {
		zap.S().Errorf("发送关闭信息失败:%s",err.Error())
		return
	}
	err = conn.Close()
	if err != nil {
		zap.S().Errorf("关闭连接失败:%s",err.Error())
		return
	}
}


