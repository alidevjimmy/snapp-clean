package http

type Rest interface {
	Start(address string) error
	ShutDown() error
}