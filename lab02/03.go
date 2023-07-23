package lab02

import (
	"fmt"
	"net"
	"time"
)

type Port struct {
	Number int
	Open   bool
}

func CheckPort(host string, port int, timeout int) bool {
	address := net.JoinHostPort(host, fmt.Sprint(port))
	conn, err := net.DialTimeout("tcp", address, time.Duration(timeout*int(time.Second)))
	if err == nil && conn != nil {
		defer conn.Close()
		return true
	}
	return false
}
