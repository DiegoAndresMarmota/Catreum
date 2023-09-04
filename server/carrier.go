package server


type RemoteDirections struct {
	RemoteDir	string
}


type RemoteRespCall struct {
	Call	RemoteDirections
	Unpack	[]byte
}


type Carrier interface {
	Sender() 							<-chan RemoteRespCall
	Transfer(Carrier) 					error
	Receiver(RemoteDirections, []byte) 	error
	Direction()							RemoteDirections
} 