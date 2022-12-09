package service

type IMudService interface {
    Connect(name string) error
    DisConnect()
    Send(command string) error
}
