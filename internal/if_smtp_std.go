package internal

import "net"

type IF_SMTP_STD interface {
	GetConn() net.Conn
	GetGreeating() string
	HandleCMD(cmd string, vars string)
}
