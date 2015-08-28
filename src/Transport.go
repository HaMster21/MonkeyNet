package MonkeyNet

type MessageRouter interface {
	FindClosestNodes(to ID, in NodeRepository, find uint) []*Node
}

type TransportMethod interface {
	Send(a Message) error
	Receive(a Message) error
}
