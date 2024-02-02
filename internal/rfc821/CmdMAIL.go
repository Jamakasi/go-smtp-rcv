package rfc821

import (
	"fmt"
	"net"
	"strings"
)

/*
MAIL (MAIL)

	This command is used to initiate a mail transaction in which
	the mail data is delivered to one or more mailboxes.  The
	argument field contains a reverse-path.

	The reverse-path consists of an optional list of hosts and
	the sender mailbox.  When the list of hosts is present, it
	is a "reverse" source route and indicates that the mail was
	relayed through each host on the list (the first host in the
	list was the most recent relay).  This list is used as a
	source route to return non-delivery notices to the sender.
	As each relay host adds itself to the beginning of the list,
	it must use its name as known in the IPCE to which it is
	relaying the mail rather than the IPCE from which the mail
	came (if they are different).  In some types of error
	reporting messages (for example, undeliverable mail
	notifications) the reverse-path may be null (see Example 7).

	This command clears the reverse-path buffer, the
	forward-path buffer, and the mail data buffer; and inserts
	the reverse-path information from this command into the
	reverse-path buffer.

S: 250 Requested mail action okay, completed
F: 552 Requested mail action aborted: exceeded storage allocation
F: 451 Requested action aborted: local error in processing
F: 452 Requested action not taken: insufficient system storage
E: 500 Syntax error, command unrecognized

	[This may include errors such as command line too long]

E: 501 Syntax error in parameters or arguments
E: 421 <domain> Service not available,

	 closing transmission channel
	[This may be a reply to any command if the service knows it
	must shut down]
*/
type CmdMAIL struct {
	connection net.Conn
	args       string
}

func NewCmdMAIL(c net.Conn, args string) *CmdMAIL {
	cmd := &CmdMAIL{
		connection: c,
		args:       args,
	}
	return cmd
}

func (cmd *CmdMAIL) RunCMD() {
	if len(cmd.args) < 5 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	cmd_test := cmd.args[:5]
	if strings.Compare(cmd_test, "FROM:") != 0 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	raw_mail := cmd.args[5:]
	fmt.Printf("MAIL \"%s\" has contain some data: \"%s\"\n", cmd_test, raw_mail)
	cmd.connection.Write([]byte("250 Requested mail action okay, completed\r\n"))
}
