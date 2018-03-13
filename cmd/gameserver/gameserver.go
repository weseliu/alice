package main

import (
	"network"
	"network/websocket"
	"proto/command"
)

func main() {
	queue := network.NewEventQueue()
	peer := network.PeerManager().New("GameServer", func() network.Peer {
		return websocket.NewAcceptor(queue)
	})
	peer.Start("http://127.0.0.1:8801/gameserver")

	peer.RegisterMessage("command.CMD_AUTH_CS", func(event *network.Event) {
		msg := event.Msg.(*command.CMD_AUTH_CS)
		msg.String()
		event.Session.Send(&command.CMD_AUTH_SC{
			RetCode: 100,
		})
	})

	queue.StartLoop()
	queue.Wait()
	peer.Stop()
}
