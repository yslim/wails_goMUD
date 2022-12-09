package socket

import (
    "net"
    "sync"

    "goMUD/pkg/log"
)

var (
    gConnections = make(map[*Connection]int)
    gMutex       sync.Mutex
)

type TcpServer struct {
    Port         int
    IncomingChan chan *TcpEvent
}

func NewTcpServer(port int, recvChan chan *TcpEvent) *TcpServer {
    return &TcpServer{Port: port, IncomingChan: recvChan}
}

func (s *TcpServer) Run() {
    laddr, err := net.ResolveTCPAddr("tcp", ":5800")
    if err != nil {
        log.Fatal("ResolveTCPAddr : error = %v", err.Error())
    }
    listener, _ := net.ListenTCP("tcp", laddr)
    for {
        tcpConn, err := listener.AcceptTCP()
        if err != nil {
            log.Error("Accept : error = %v", err.Error())
            continue
        }

        NewConnection(tcpConn, s.IncomingChan)
    }
}
