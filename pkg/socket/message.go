package socket

type MsgType int

const (
	Connected MsgType = iota + 1
	Closed
	Messgae
)

type TcpEvent struct {
	MType   MsgType
	Client  *Client
	Message []byte
}
