package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"net/http"
	"ws_srv/global"
	"ws_srv/proto/gen/go/msgpb"
	"ws_srv/utils"
)
var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
type UserConn struct{
	UserClaims *utils.CustomClaims
	UserConn *websocket.Conn
}

var UserClientConn = make(map[uint32]*UserConn)

var test = 0

func Test(c *gin.Context) {
	defer func() {
		fmt.Println("协程关闭....")
	}()

	if test == 0 {
		//go func() {
		//	for {
		//		fmt.Println(runtime.NumGoroutine(),test)
		//		time.Sleep(time.Second * 3)
		//	}
		//}()
	}

	test++
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Errorf("初始websocket失败:%s", err.Error())
		return
	}

	for {
		//接受消息
		messageType, receive, err := conn.ReadMessage()
		if messageType < 0 {
			break
		}

		if messageType != websocket.BinaryMessage {
			continue
		}
		if err != nil {
			zap.S().Errorf("读取消息失败:%s",err.Error())
			continue
		}

		msg := &msgpb.Msg{}
		err = proto.Unmarshal(receive,msg)
		if err != nil {
			zap.S().Errorf("解析消息失败:%s",err.Error())
			continue
		}
		handleFunc,ok := global.SocketRouter[msg.GetPath()]
		if !ok {
			zap.S().Error("消息路由不存在")
			continue
		}
		err = handleFunc(conn,msg)
		if err != nil {
			break
		}
		fmt.Printf("%#v\n",UserClientConn)
	}
}
