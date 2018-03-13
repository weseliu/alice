package network

import (
	"sync"
)

type Peer interface {
	SetName(string)
	Name() string

	SetAddress(string)
	Address() string

	SetTag(interface{})
	Tag() interface{}

	IsRunning() bool
	SetRunning(v bool)

	Start(address string) Peer
	Stop()

	Encoder
	MessageHandler
}

type PeerBase struct {
	name    string
	address string
	tag     interface{}

	running      bool
	runningGuard sync.RWMutex

	EventQueue
	SessionManager
	Encoder
	MessageHandler
}

func (Self *PeerBase) SetName(name string) {
	Self.name = name
}

func (Self *PeerBase) Name() string {
	return Self.name
}

func (Self *PeerBase) SetAddress(address string) {
	Self.address = address
}

func (Self *PeerBase) Address() string {
	return Self.address
}

func (Self *PeerBase) SetTag(tag interface{}) {
	Self.tag = tag
}

func (Self *PeerBase) Tag() interface{} {
	return Self.tag
}

func (Self *PeerBase) IsRunning() bool {
	Self.runningGuard.RLock()
	defer Self.runningGuard.RUnlock()

	return Self.running
}

func (Self *PeerBase) SetRunning(v bool) {
	Self.runningGuard.Lock()
	Self.running = v
	Self.runningGuard.Unlock()
}

func (Self *PeerBase) Queue() EventQueue {
	return Self.EventQueue
}

func (Self *PeerBase) Start(address string) Peer{
	return nil
}

func (Self *PeerBase) Stop(){

}
