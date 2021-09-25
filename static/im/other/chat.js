function Chat(
    {
        socket,
    }
) {
    this.opts = {
    };
    this.socket = socket;//websocket实例
}

Chat.prototype.connect = function (token) {
    let msg = new proto.Msg.Msg();
    msg.setMsgType(proto.Msg.MsgType.AUTH)
    msg.setPath("/auth")
    msg.setToken(token)

    let binary = msg.serializeBinary()
    this.socket.send(binary)
};

if (typeof window != 'undefined') window.Chat = Chat;