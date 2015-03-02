package connection

type MessageType int8

const (
	SystemMsg MessageType = 1 << iota
	PingMsg
	DataMsg
)

type Message struct {
	Type MessageType
	Data interface{}
}

func NewMessage() *Message {
	return &Message{}
}
