package service

import "fmt"

type MudService struct {
}

var mudService = &MudService{}

func GetMudService() *MudService {
    return mudService
}

func (m *MudService) Connect(name string) error {
    return nil
}

func (m *MudService) DisConnect() {
}

func (m *MudService) Send(command string) error {
    fmt.Printf("command = %s\n", command)
    return nil
}
