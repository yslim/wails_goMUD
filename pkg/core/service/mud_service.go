package service

import (
    "fmt"
    "github.com/wailsapp/wails/v2/pkg/runtime"
    "github.com/yslim/go-util"
    "goMUD/pkg/app"
    "goMUD/pkg/config"
    "goMUD/pkg/log"
    "goMUD/pkg/socket"
    "strings"
)

type MudService struct {
    *socket.Connection
}

var mudService = &MudService{nil}

func GetMudService() *MudService {
    return mudService
}

func (m *MudService) Connect(name string) error {
    tcpConn, err := socket.ConnectTo(config.Mud.Server, config.Mud.Port)
    if err != nil {
        return err
    }

    tcpMesgChannel := make(chan *socket.TcpEvent, 100)

    m.Connection = socket.NewConnection(tcpConn, tcpMesgChannel)

    go onTcpEvent(tcpMesgChannel)

    return nil
}

func (m *MudService) DisConnect() {
    m.Close()
    m.Connection = nil
}

func (m *MudService) Send(command string) error {
    fmt.Printf("command = %s\n", command)
    bytes, err := util.FromUTF8("euckr", []byte(command+"\n"))
    if err != nil {
        log.Error("util.FromUTF8 error = ", err)
        return err
    }

    m.Connection.SendMessage(bytes)
    return nil
}

func onTcpEvent(incomingChan chan *socket.TcpEvent) {
    for {
        tcpEvent := <-incomingChan
        if tcpEvent.MType == socket.Connected {
            log.Info("TCP Connected, RemoteAddr=%v", string(tcpEvent.Message))
        } else if tcpEvent.MType == socket.Closed {
            log.Info("TCP Closed, reason=%v", string(tcpEvent.Message))
        } else {
            onMessageFromTcp(tcpEvent.Conn, tcpEvent.Message)
        }
    }
}

func onMessageFromTcp(conn *socket.Connection, bytes []byte) {
    utfStr, _ := util.ToUTF8("euc-kr", bytes)
    message := strings.TrimSpace(string(utfStr)) + "\n\r"

    fmt.Printf("message = %s", message)

    runtime.EventsEmit(app.App().Ctx(), "OnMessage", message)
}
