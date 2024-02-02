package rfc821

import (
	"net"
)

/*
HELLO (HELO)

	This command is used to identify the sender-SMTP to the
	receiver-SMTP.  The argument field contains the host name of
	the sender-SMTP.

	The receiver-SMTP identifies itself to the sender-SMTP in
	the connection greeting reply, and in the response to this
	command.

	This command and an OK reply to it confirm that both the
	sender-SMTP and the receiver-SMTP are in the initial state,
	that is, there is no transaction in progress and all state
	tables and buffers are cleared.

HELO <SP> <domain> <CRLF>

Return codes (4.2.1):
S: 250 Requested mail action okay, completed
E: 500 Syntax error, command unrecognized

	[This may include errors such as command line too long]

E: 501 Syntax error in parameters or arguments
E: 504 Command parameter not implemented
E: 421 <domain> Service not available,

	 closing transmission channel
	[This may be a reply to any command if the service knows it
	must shut down]
*/
type CmdHELO struct {
	connection net.Conn
	args       string
}

func NewCmdHELO(c net.Conn, args string) *CmdHELO {
	cmd := &CmdHELO{
		connection: c,
		args:       args,
	}
	return cmd
}

func (cmd *CmdHELO) RunCMD() {
	if len(cmd.args) == 0 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	cmd.connection.Write([]byte("250 Requested mail action okay, completed\r\n"))
}
