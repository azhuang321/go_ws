package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/httphead"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"io"
	"log"
	"net"
	"net/http"
	"runtime"
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

func Test1(c *gin.Context) {
	go func() {
		ln, err := net.Listen("tcp", "0.0.0.0:9003")
		if err != nil {
			panic(err)
		}

		// Prepare handshake header writer from http.Header mapping.
		header := ws.HandshakeHeaderHTTP(http.Header{
			"X-Go-Version": []string{runtime.Version()},
		})

		u := ws.Upgrader{
			OnHost: func(host []byte) error {
				//if string(host) == "github.com" {
				//	return nil
				//}
				//return ws.RejectConnectionError(
				//	ws.RejectionStatus(403),
				//	ws.RejectionHeader(ws.HandshakeHeaderString(
				//		"X-Want-Host: github.com\r\n",
				//	)),
				//)
				return nil
			},
			OnHeader: func(key, value []byte) error {
				if string(key) != "Cookie" {
					return nil
				}
				ok := httphead.ScanCookie(value, func(key, value []byte) bool {
					// Check session here or do some other stuff with cookies.
					// Maybe copy some values for future use.
					return true
				})
				if ok {
					return nil
				}
				return ws.RejectConnectionError(
					ws.RejectionReason("bad cookie"),
					ws.RejectionStatus(400),
				)
			},
			OnBeforeUpgrade: func() (ws.HandshakeHeader, error) {
				return header, nil
			},
		}
		for {
			conn, err := ln.Accept()
			if err != nil {
				log.Fatal(err)
			}
			_, err = u.Upgrade(conn)
			if err != nil {
				log.Printf("upgrade error: %s", err)
			}
			var (
				state  = ws.StateServerSide
				reader = wsutil.NewReader(conn, state)
				writer = wsutil.NewWriter(conn, state, ws.OpText)
			)
			go func() {
				for {
					header, err := reader.NextFrame()
					//if header.OpCode == ws.OpClose {
					//	fmt.Println(err)
					//
					//	return
					//}
					if err != nil {
						fmt.Println(err)
						conn.Close()
						break
					}

					// Reset writer to write frame with right operation code.
					writer.Reset(conn, state, header.OpCode)
					//err = ws.RejectConnectionError(
					//	ws.RejectionReason("bad cookie"),
					//	ws.RejectionStatus(400),
					//)
					//conn.
					fmt.Println(err)
					conn.Close()

					if _, err = io.Copy(writer, reader); err != nil {
						// handle error
					}
					if err = writer.Flush(); err != nil {
						// handle error
					}
				}
			}()
		}
	}()
}

func Test2(ctx *gin.Context)  {

}


