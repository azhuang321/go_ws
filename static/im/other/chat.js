function Chat(
    {
        loginUserInfo,
        layim,
        $,
    }
) {
    this.opts = {
    };
    this.socket = {};
    this.$ = $;
    this.layim = layim;

    this.initWebsocket(loginUserInfo)
}

Chat.prototype.initWebsocket = function (loginUserInfo) {
    //设置心跳信息
    // let pingMsgObj = new proto.Msg.HeartMsg()
    // pingMsgObj.setPing("ping")
    // let pingBinary = pingMsgObj.serializeBinary()

    const url = 'ws://wslhost:9002/v1/websocket?token=' + loginUserInfo.token;
    let firstHeartbeat = true;
    this.socket = new WebsocketHeartbeatJs({
        url: url,
        pingTimeout: 8000,
        pongTimeout: 8000,
        // pingMsg: pingBinary
        pingMsg: "ping"
    });

    this.socket.setBinaryType()
    console.log(`连接到: ${url}`);

    let mySocket = this.socket;
    this.socket.onopen = function () {
        console.log('连接成功...');
        setTimeout(() => {
            console.dir(`等待心跳 ${mySocket.opts.pingTimeout} ms 将会有心跳消息: '${mySocket.opts.pingMsg}' 到达`);
        }, 1500);
    }
    let layim = this.layim
    this.socket.onmessage = function (e) {
        let receiveMsg = proto.Msg.Msg.deserializeBinary(e.data)
        let receiveMsgInfo = receiveMsg.getContent().getReceiveInfo()

        layim.getMessage({
            username: receiveMsg.getContent().getSendInfo().getSendUserInfo().getUsername()
            , avatar: receiveMsgInfo.getReceiveUserInfo().getAvatar()
            , id: receiveMsg.getContent().getSendInfo().getSendUserInfo().getId()
            , type: "friend"
            , cid: 0 //模拟消息id，会赋值在li的data-cid上，以便完成一些消息的操作（如撤回），可不填
            , content: receiveMsgInfo.getContent()
        });


        // if (e.data == websocketHeartbeatJs.opts.pingMsg && firstHeartbeat) {
        //     setTimeout(() => {
        //         console.dir(`Close your network, wait ${websocketHeartbeatJs.opts.pingTimeout + websocketHeartbeatJs.opts.pongTimeout}+ ms, websocket will reconnect`, 'cadetblue');
        //     }, 1500);
        //     firstHeartbeat = false;
        // }
    }
    this.socket.onreconnect = function () {
        console.log(`断线,正在重连中...`);
    }
    this.socket.onclose = function (event) {
        //todo 实现重新连接
        console.log(`服务器主动关闭:` + event.reason);
        layer.confirm('已断线是否重新连接', {icon: 2, title:'提示'}, function(index){
            layer.load(1);
            layer.close(index);
        });
    }
};

Chat.prototype.getUserInfo = function (loginUserInfo) {
    let userInfo = {}
    this.$.ajax({url:"http://wslhost:9001/v1/user/info",
        method:"POST",
        dataType:"json",
        data: JSON.stringify({mobile:loginUserInfo.mobile}),
        async:false,
        beforeSend: function (request){
            request.setRequestHeader("x-token",loginUserInfo.token)
        },
        success:function(result){
            if (result.code == 0) {
                userInfo = result.data
            }else{
                layer.msg(result.msg,{icon:2})
            }
        },
        error:function(error){
            layer.msg(error.responseJSON.msg,{icon:2})
        }
    });
    userInfo.username = userInfo.mobile
    return userInfo;
};

Chat.prototype.getFriendList = function (loginUserInfo) {
    let friendList = {}
    this.$.ajax({url:"http://wslhost:9003/v1/chat/getUserFriendList",
        method:"POST",
        async:false,
        beforeSend: function (request){
            request.setRequestHeader("x-token",loginUserInfo.token)
        },
        success:function(result){
            if (result.code == 0) {
                friendList = result.data
            }else{
                layer.msg(result.msg,{icon:2})
            }
        },
        error:function(error){
            layer.msg(error.responseJSON.msg,{icon:2})
        }
    });
    return friendList;
};


if (typeof window != 'undefined') window.Chat = Chat;