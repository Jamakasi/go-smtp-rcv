package rfc821

import "net"

/*
S: 221 <domain> Service closing transmission channel
E: 500
*/
type CmdQUIT struct {
	connection net.Conn
}

func NewCmdQUIT(c net.Conn) *CmdQUIT {
	cmd := &CmdQUIT{
		connection: c,
	}
	return cmd
}

func (cmd *CmdQUIT) RunCMD() {
	cmd.connection.Write([]byte("221 <domain> Service closing transmission channel\r\n"))
	cmd.connection.Close()
}
