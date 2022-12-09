package socket

import (
    "fmt"
    "goMUD/pkg/log"
    "net"
)

func ConnectTo(ip string, port uint16) (*net.TCPConn, error) {
    raddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", ip, port))
    if err != nil {
        log.Error("ResolveTCPAddr : error = %v", err.Error())
        return nil, err
    }

    tcpConn, err := net.DialTCP("tcp", nil, raddr)
    if err != nil {
        log.Error("net.DialTCP, error = %v", err)
        return nil, err
    }

    return tcpConn, nil
}
