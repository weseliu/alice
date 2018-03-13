package network

type peerManager struct {
	peerMap map[string] Peer
}

var instance *peerManager
func PeerManager() *peerManager{
	if instance == nil{
		instance = &peerManager{}
		instance.peerMap = make(map[string] Peer)
	}
	return instance
}

func (Self *peerManager) New(name string, creator func() Peer) Peer{
	if _, ok := Self.peerMap[name]; !ok {
		Self.peerMap[name] = creator()
		Self.peerMap[name].SetName(name)
	}
	return Self.peerMap[name]
}

func (Self *peerManager) Get(name string) Peer{
	if peer, ok := Self.peerMap[name]; ok {
		return peer
	}
	return nil
}
