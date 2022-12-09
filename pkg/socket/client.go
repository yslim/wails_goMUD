package socket

import (
   "bufio"
   "net"

   "goMUD/pkg/config"
   "goMUD/pkg/log"
)

type Client struct {
   incoming chan *TcpEvent
   outgoing chan []byte
   reader   *bufio.Reader
   writer   *bufio.Writer
   conn     net.Conn
}

func (client *Client) SendMessage(bytes []byte) {
   client.outgoing <- bytes
}

func (client *Client) listen() {
   go client.read()
   go client.write()
}

func (client *Client) read() {
   for {
      line, err := client.reader.ReadBytes('\n')
      if err != nil {
         client.Close()
         client.incoming <- &TcpEvent{MType: Closed, Client: client, Message: []byte(err.Error())}
         break
      } else {
         client.incoming <- &TcpEvent{MType: Messgae, Client: client, Message: line}
      }
   }
}

func (client *Client) write() {
   for data := range client.outgoing {
      _, err := client.writer.Write(data)
      if err != nil {
         // socket error check & closing 은 read 부분에서 수행한다.
         log.Error("socket write error = %v", err.Error())
         continue
      }
      _ = client.writer.Flush()
   }
}

func NewClient(connection *net.TCPConn, callbackChan chan *TcpEvent) *Client {
   _ = connection.SetReadBuffer(config.Socket.BufferSize)
   _ = connection.SetNoDelay(true)

   writer := bufio.NewWriter(connection)
   reader := bufio.NewReader(connection)

   client := &Client{
      incoming: callbackChan,
      outgoing: make(chan []byte, 100),
      conn:     connection,
      reader:   reader,
      writer:   writer,
   }

   client.incoming <- &TcpEvent{MType: Connected, Client: client, Message: []byte(connection.RemoteAddr().String())}

   client.listen()

   return client
}

func (client *Client) Close() {
   _ = client.conn.Close()
   gMutex.Lock()
   delete(gClients, client)
   gMutex.Unlock()
}

func (client *Client) RemoteAddr() string {
   return client.conn.RemoteAddr().String()
}
