package server

import (
	"fmt"
	"sync"
)


type CallPacket struct {
	direction 	RemoteDirections
	network     map[RemoteDirections]*CallPacket
	channel   	chan RemoteRespCall
	pack   		sync.RWMutex
}


func NewCallPacket(direction RemoteDirections) *CallPacket {
	return &CallPacket{
		direction: 	direction,
		channel: 	make(chan RemoteRespCall, 3000),
		network: 	make(map[RemoteDirections]*CallPacket),
	}
}


func (c *CallPacket) Channel() <-chan RemoteRespCall {
	return c.channel
}


func (c *CallPacket) Networking(ntw Carrier) error {
	c.pack.Lock()
	defer c.pack.Unlock()
	c.network[ntw.Direction()] = ntw
	return nil
}


func (c *CallPacket) Direction() RemoteDirections {
	return c.direction
}


func (c *CallPacket) TransferMessage(to RemoteDirections, Unpack []byte) error {
	c.pack.RLock()
	defer c.pack.RUnlock()

	pack, ok := c.network[to]
	if !ok {
		return fmt.Errorf("no such direction %s", to)
	}

	pack.channel <- RemoteRespCall{
		Call: 		c.direction,
		Unpack: 	Unpack,

	}

	return nil
}

