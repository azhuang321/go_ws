package api

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"os"
	"time"
	"ws_srv/global"
	"ws_srv/logic"
	"ws_srv/proto/gen/go/msgpb"
)

var test = false

func WebSocket(ctx *gin.Context) {

	if !test {
		go testConsumer()
	}






	userId,_ := ctx.Get("userId")
	currentUserId := userId.(uint32)
	conn, err := logic.UpGrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		zap.S().Errorf("初始websocket失败:%s", err.Error())
		return
	}
	defer logic.ConnClose(conn,currentUserId)
	logic.UserClientConn.Store(currentUserId,conn)


	//发送消息到消息队列
	rlog.SetLogLevel("error")
	p, _ := rocketmq.NewProducer(
		producer.WithGroupName("testGroup"),
		producer.WithNsResolver(primitive.NewPassthroughResolver([]string{"172.18.0.1:9876"})),
		producer.WithCreateTopicKey("test"),
		producer.WithRetry(2),
	)
	err = p.Start()
	if err != nil {
		fmt.Printf("start producer error: %s\n", err.Error())
		os.Exit(1)
	}

	for {
		//conn.SetCloseHandler()
		//接受消息
		//messageType, receive, err := conn.ReadMessage()
		_, receive, err := conn.ReadMessage()
		if err != nil {
			closeErr := err.(*websocket.CloseError)
			zap.S().Errorf("接收消息失败:%d:%s",closeErr.Code,closeErr.Error())
			break
		}

		//msg,err := logic.CheckAndUnmarshalMsg(messageType,receive)
		//if err != nil {
		//	logic.CloseConn(currentUserId,"解析消息失败")
		//	break
		//}


		rmqMsg := &primitive.Message{
			Topic: "test",
			Body:  receive,
		}

		res, err := p.SendSync(context.Background(), rmqMsg)

		if err != nil {
			fmt.Printf("send message error: %s\n", err)
		} else {
			fmt.Printf("send message success: result=%s\n", res.String())
		}
	}
	err = p.Shutdown()
	if err != nil {
		fmt.Printf("shutdown producer error: %s\n", err.Error())
	}
}



func testConsumer(){
	test = true
	rlog.SetLogLevel("error")
	c, _ := rocketmq.NewPushConsumer(
		consumer.WithGroupName("testGroup"),
		consumer.WithNsResolver(primitive.NewPassthroughResolver([]string{"172.18.0.1:9876"})),
	)
	err := c.Subscribe("test", consumer.MessageSelector{}, func(ctx context.Context, msgs ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range msgs {
			msg := &msgpb.Msg{}
			msg,err := logic.CheckAndUnmarshalMsg(websocket.BinaryMessage,msgs[i].Body)

			if err != nil {
				return consumer.ConsumeRetryLater, err
			}
			handleFunc,ok := global.SocketRouter[msg.GetPath()]
			if !ok {
				zap.S().Error("消息路由不存在")
				logic.CloseConn(1,"消息错误")
				return consumer.ConsumeRetryLater, err
			}

			err = handleFunc(nil,msg)
			if err != nil {
				logic.CloseConn(1,"消息处理失败")
				return consumer.ConsumeRetryLater, err

			}

		}

		return consumer.ConsumeSuccess, nil
	})
	fmt.Println("end.......")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Note: start after subscribe
	err = c.Start()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	time.Sleep(time.Hour)
	err = c.Shutdown()
	if err != nil {
		fmt.Printf("shutdown Consumer error: %s", err.Error())
	}
}

