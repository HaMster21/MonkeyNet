package MonkeyNet

type KBucket interface {
	Add(a Node)
	ClosestNode(to uint64) Node
}
