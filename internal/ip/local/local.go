package local

import (
	"net"
	"nvidia_gpu_exporter/internal/ip"
)

type ipLocal struct {
}

func (i *ipLocal) Get() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP.String(), nil
}

func New() ip.IP {
	return &ipLocal{}
}
