package api

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"sync"
	"ws_srv/proto/gen/go/msgpb"
)

var upGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var ClientList = make(map[string]*User)

type User struct {
	Conn *websocket.Conn
	Id   string
	Mux  sync.Mutex
}
type TypeMessage struct {
	Type interface{} `json:"type"`
	Data interface{} `json:"data"`
}

var KefuList = make(map[string][]*User)

func Test(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		zap.S().Errorf("初始websocket失败:%s", err.Error())
		return
	}
	user := &User{
		Conn: conn,
		Id:   "1",
	}
	// 挤下线
	oldUser, ok := ClientList[user.Id]
	if oldUser != nil || ok {
		msg := TypeMessage{
			Type: "close",
			Data: user.Id,
		}
		str, _ := json.Marshal(msg)
		if err := oldUser.Conn.WriteMessage(websocket.TextMessage, str); err != nil {
			oldUser.Conn.Close()
			delete(ClientList, user.Id)
		}
	}

	ClientList[user.Id] = user

	go WsServerBackend()

	for {
		//接受消息
		messageType, receive, err := conn.ReadMessage()
		fmt.Println("---------------------------------------------")
		fmt.Println(messageType)
		testMsg := &msgpb.Msg{}
		proto.Unmarshal(receive, testMsg)
		fmt.Println("*********************************************")

		if err != nil {
			for _, visitor := range ClientList {
				if visitor.Conn == conn {
					log.Println("删除用户", visitor.Id)
					delete(ClientList, visitor.Id)
					//VisitorOffline(visitor.To_id, visitor.Id, visitor.Name)
				}
			}
			log.Println(err)
			return
		}

		message <- &Message{
			conn:        conn,
			content:     receive,
			context:     c,
			messageType: messageType,
		}
	}

}

type Message struct {
	conn        *websocket.Conn
	context     *gin.Context
	content     []byte
	messageType int
	Mux         sync.Mutex
}

var message = make(chan *Message, 10)

func WsServerBackend() {
	for {
		message := <-message
		var typeMsg TypeMessage
		json.Unmarshal(message.content, &typeMsg)

		typeMsg.Type = "inputing"

		conn := message.conn
		conn.WriteMessage(websocket.TextMessage, []byte("12323"))

		if typeMsg.Type == nil || typeMsg.Data == nil {
			continue
		}
		msgType := typeMsg.Type.(string)
		log.Println("客户端:", string(message.content))

		switch msgType {
		//心跳
		case "ping":
			msg := TypeMessage{
				Type: "pong",
			}
			str, _ := json.Marshal(msg)
			message.Mux.Lock()
			defer message.Mux.Unlock()
			conn.WriteMessage(websocket.TextMessage, str)
		case "inputing":
			data := typeMsg.Data.(map[string]interface{})
			//from := data["from"].(string)
			to := data["to"].(string)
			fmt.Printf("%+v\n", to)

		}

	}

}
