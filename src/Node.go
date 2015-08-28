package MonkeyNet

type InfoHash struct {
	Value uint64
}

type Node interface {
	ID() InfoHash
}

type NodeRepository interface {
	Add(a Node, with TransportMethod)
	GetTransportMethods(fora Node) []TransportMethod
}
