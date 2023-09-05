package server

import (
	"time"
	"fmt"
)

type Net struct {
	directCh 	chan RemoteRespCall
	quitCh		chan struct{} 
	NetOptions 		
}


type NetOptions struct {
	Carriers []Carrier
}


func NewNet(options NetOptions) *Net {
	return &Net{
		NetOptions: 	options,
		directCh: 		make(chan RemoteRespCall),
		quitCh: 		make(chan struct{}, 1),
	}
}


func (n *Net) Init() {
	n.InitCarriers()
	token := time.NewTicker(10 * time.Second)

	init:
		for {
		select {
		case chn := <- n.directCh:
			fmt.Println(chn)
		case <- n.quitCh: 
			break init
		case <- token.C:
			fmt.Println("Initialize...")
		}
	}

}


func (n *Net) InitCarriers() {
	for _, nt := range n.NetOptions.Carriers {
		go func(nt Carrier) {
			for chn := range nt.Sender() {
				n.directCh <- chn
			}
		} (nt)
	}
}