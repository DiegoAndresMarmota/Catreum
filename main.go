package main

import (
	"time"
	"github.com/DiegoAndresMarmota/Catreum/server"
)

func main() {
	ntLocal := server.CallPacket("LOCAL")
	ntRemote := server.CallPacket("REMOTE")

	ntLocal.Networking(ntRemote)
	ntRemote.Networking(ntLocal)

	go func() {
		for {
			ntRemote.Receiver(ntLocal.Direction(), []byte("Running"))
			time.Sleep(3 * time.Second)
		}
	}()

	options := server.NetOptions {
		Carriers: []server.Carrier{ntLocal},
	}

	n := server.NewNet(options)
	n.Init()
}