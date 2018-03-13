package network

import (
	"fmt"
	"sync/atomic"
)

type EventType int32
const (
	EventNone          EventType = iota
	EventConnected
	EventConnectFailed
	EventAccepted
	EventAcceptFailed
	EventClosed
	EventReceive
	EventSend
)

func (Self EventType) String() string {
	switch Self {
	case EventNone:
		return "None"
	case EventConnected:
		return "Connect"
	case EventConnectFailed:
		return "ConnectFailed"
	case EventAccepted:
		return "Accepted"
	case EventAcceptFailed:
		return "AcceptFailed"
	case EventClosed:
		return "Closed"
	case EventReceive:
		return "Receive"
	case EventSend:
		return "Send"
	}
	return fmt.Sprintf("unknown(%d)", Self)
}

type EventResult int32
const (
	EventResultOK            EventResult = iota
	EventResultSocketError
	EventResultSocketTimeout
	EventResultPackageCrack
	EventResultCodecError
	EventResultRequestClose
	EventResultNextChain
)

type Event struct {
	UID     int64
	Type    EventType
	MsgID   uint32
	Msg     interface{}
	Data    []byte
	Tag     interface{}
	Session Session
}

func (Self *Event) Clone() *Event {
	c := &Event{
		UID:        Self.UID,
		Type:       Self.Type,
		MsgID:      Self.MsgID,
		Msg:        Self.Msg,
		Tag:        Self.Tag,
		Session:    Self.Session,
		Data:       make([]byte, len(Self.Data)),
	}
	copy(c.Data, Self.Data)
	return c
}

var eventUid int64
func genEventUID() int64 {
	atomic.AddInt64(&eventUid, 1)
	return eventUid
}

func NewEvent(t EventType, s Session) *Event {
	self := &Event{
		Type: t,
		Session:  s,
		UID : genEventUID(),
	}
	return self
}


