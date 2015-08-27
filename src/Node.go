package MonkeyNet

import (
	"hash"
	"io"
)

type Node interface {
	ID() hash.Hash64
	io.Reader
	io.Writer
}
