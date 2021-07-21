package core

type Service interface {
	Hostname() string
	Address() string
	Port() uint16
}
