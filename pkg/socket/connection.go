package socket

import (
    "bufio"
    "net"

    "goMUD/pkg/config"
    "goMUD/pkg/log"
)

type Connection struct {
    incoming chan *TcpEvent
    outgoing chan []byte
    reader   *bufio.Reader
    writer   *bufio.Writer
    conn     net.Conn
}

func (conn *Connection) SendMessage(bytes []byte) {
    conn.outgoing <- bytes
}

func (conn *Connection) listen() {
    go conn.read()
    go conn.write()
}

func (conn *Connection) read() {
    for {
        line, err := conn.reader.ReadBytes('\n')
        if err != nil {
            conn.Close()
            conn.incoming <- &TcpEvent{MType: Closed, Conn: conn, Message: []byte(err.Error())}
            break
        } else {
            conn.incoming <- &TcpEvent{MType: Messgae, Conn: conn, Message: line}
        }
    }
}

func (conn *Connection) write() {
    for data := range conn.outgoing {
        _, err := conn.writer.Write(data)
        if err != nil {
            // socket error check & closing 은 read 부분에서 수행한다.
            log.Error("socket write error = %v", err.Error())
            continue
        }
        _ = conn.writer.Flush()
    }
}

func NewConnection(tcpConn *net.TCPConn, callbackChan chan *TcpEvent) *Connection {
    _ = tcpConn.SetReadBuffer(config.Socket.BufferSize)
    _ = tcpConn.SetNoDelay(true)

    writer := bufio.NewWriter(tcpConn)
    reader := bufio.NewReader(tcpConn)

    conn := &Connection{
        incoming: callbackChan,
        outgoing: make(chan []byte, 100),
        conn:     tcpConn,
        reader:   reader,
        writer:   writer,
    }

    gMutex.Lock()
    gConnections[conn] = 1
    gMutex.Unlock()

    conn.incoming <- &TcpEvent{MType: Connected, Conn: conn, Message: []byte(tcpConn.RemoteAddr().String())}

    conn.listen()

    return conn
}

func (conn *Connection) Close() {
    _ = conn.conn.Close()
    gMutex.Lock()
    delete(gConnections, conn)
    gMutex.Unlock()
}

func (conn *Connection) RemoteAddr() string {
    return conn.conn.RemoteAddr().String()
}
