package MonkeyNet

import "bytes"

type Message interface {
	Verb() string
	Version() string
	Payload() bytes.Buffer
}
