package network

type Session interface {
	SetID(id int64)
	ID() int64

	SetTag(tag interface{})
	Tag() interface{}

	Peer() Peer
	Send(interface{})
	Close()
}


