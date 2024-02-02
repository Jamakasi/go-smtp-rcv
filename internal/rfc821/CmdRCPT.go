package rfc821

import (
	"fmt"
	"net"
	"strings"
)

/*
RECIPIENT (RCPT)

            This command is used to identify an individual recipient of
            the mail data; multiple recipients are specified by multiple
            use of this command.

            The forward-path consists of an optional list of hosts and a
            required destination mailbox.  When the list of hosts is
            present, it is a source route and indicates that the mail
            must be relayed to the next host on the list.  If the
            receiver-SMTP does not implement the relay function it may
            user the same reply it would for an unknown local user
            (550).

            When mail is relayed, the relay host must remove itself from
            the beginning forward-path and put itself at the beginning
            of the reverse-path.  When mail reaches its ultimate
            destination (the forward-path contains only a destination
            mailbox), the receiver-SMTP inserts it into the destination
            mailbox in accordance with its host mail conventions.
RCPT
S: 250, 251
F: 550, 551, 552, 553, 450, 451, 452
E: 500, 501, 503, 421
*/

type CmdRCPT struct {
	connection net.Conn
	args       string
}

func NewCmdRCPT(c net.Conn, args string) *CmdRCPT {
	cmd := &CmdRCPT{
		connection: c,
		args:       args,
	}
	return cmd
}

func (cmd *CmdRCPT) RunCMD() {
	if len(cmd.args) < 5 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	cmd_test := cmd.args[:3]
	if strings.Compare(cmd_test, "TO:") != 0 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	raw_mail := cmd.args[3:]
	fmt.Printf("RCPT \"%s\" has contain some data: \"%s\"\n", cmd_test, raw_mail)
	cmd.connection.Write([]byte("250 Requested mail action okay, completed\r\n"))
}
