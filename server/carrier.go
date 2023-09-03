package server


type RemoteDirections struct {
	RemoteDir	string
}


type RemoteRespCall struct {
	Call	RemoteDirections
	Unpack	[]byte
}


type Carrier interface {
	Return() 							<-chan RemoteRespCall
	Message(Carrier) 					error
	Receive(RemoteDirections, []byte) 	error
	Direction()							RemoteDirections
} 