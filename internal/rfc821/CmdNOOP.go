package rfc821

import "net"

/*
RFC 821
NOOP (NOOP)

	This command does not affect any parameters or previously
	entered commands.  It specifies no action other than that
	the receiver send an OK reply.

	This command has no effect on any of the reverse-path
	buffer, the forward-path buffer, or the mail data buffer.

NOOP <CRLF>

Return codes (4.2.1):
S: 250 Requested mail action okay, completed
E: 500 Syntax error, command unrecognized

	[This may include errors such as command line too long]

E: 421 <domain> Service not available,

	 closing transmission channel
	[This may be a reply to any command if the service knows it
	must shut down]
*/
type CmdNOOP struct {
	connection net.Conn
}

func NewCmdNOOP(c net.Conn) *CmdNOOP {
	cmd := &CmdNOOP{
		connection: c,
	}
	return cmd
}

func (cmd *CmdNOOP) RunCMD() {
	cmd.connection.Write([]byte("250 Requested mail action okay, completed\r\n"))
}
