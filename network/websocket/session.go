package websocket

import (
	"github.com/weseliu/alice/network"
	"github.com/gorilla/websocket"
)

type wsSession struct {
	id int64
	peer network.Peer
	conn *websocket.Conn
	tag interface{}
	OnClose func()
}

func (Self *wsSession) SetID(id int64){
	Self.id = id
}

func (Self *wsSession) ID() int64{
	return Self.id
}

func (Self *wsSession) SetTag(tag interface{}){
	Self.tag = tag
}

func (Self *wsSession) Tag() interface{}{
	return Self.tag
}

func (Self *wsSession) Peer() network.Peer{
	return Self.peer
}

func (Self *wsSession) Send(data interface{}){
	event := network.NewEvent(network.EventSend, Self)
	event.Msg = data
	Self.RawSend(event)
}

func (Self *wsSession) RawSend(event *network.Event)  {
	event.Session = Self

	go func() {
		raw := Self.Peer().Encode(event.Msg)
		Self.conn.WriteMessage(websocket.BinaryMessage, raw)
	}()
}

func (Self *wsSession) Close(){
	Self.conn.Close()
}

func (Self *wsSession) readPacket() (msgId uint32, msg interface{}, result network.EventResult) {
	t, raw, err := Self.conn.ReadMessage()
	if err != nil {
		return 0, nil, network.EventResultSocketError
	}

	switch t {
	case websocket.TextMessage:
		return 0, nil, network.EventResultCodecError
	case websocket.CloseMessage:
		return 0, nil, network.EventResultRequestClose
	case websocket.BinaryMessage:
		msgId, msg = Self.Peer().Decode(raw)
	}

	return msgId, msg, network.EventResultOK
}

func (Self *wsSession) receiveThread() {
	for {
		msgId, msg, result := Self.readPacket()
		if result == network.EventResultOK {
			ev := network.NewEvent(network.EventReceive, Self)
			ev.MsgID = msgId
			ev.Msg = msg

			var acceptor = Self.Peer().(*wsAcceptor)
			acceptor.Queue().Post(func(){
				acceptor.OnEvent(ev)
			})
		}
	}
}

func (Self *wsSession) run() {
	go Self.receiveThread()
}

func NewSession(c *websocket.Conn, p network.Peer) *wsSession {
	self := &wsSession{
		peer: p,
		conn: c,
	}
	return self
}

