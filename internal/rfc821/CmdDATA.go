package rfc821

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/textproto"
)

/*
DATA (DATA)

	The receiver treats the lines following the command as mail
	data from the sender.  This command causes the mail data
	from this command to be appended to the mail data buffer.
	The mail data may contain any of the 128 ASCII character
	codes.

	The mail data is terminated by a line containing only a
	period, that is the character sequence "<CRLF>.<CRLF>" (see
	Section 4.5.2 on Transparency).  This is the end of mail
	data indication.

	The end of mail data indication requires that the receiver
	must now process the stored mail transaction information.
	This processing consumes the information in the reverse-path
	buffer, the forward-path buffer, and the mail data buffer,
	and on the completion of this command these buffers are
	cleared.  If the processing is successful the receiver must
	send an OK reply.  If the processing fails completely the
	receiver must send a failure reply.

	When the receiver-SMTP accepts a message either for relaying
	or for final delivery it inserts at the beginning of the
	mail data a time stamp line.  The time stamp line indicates
	the identity of the host that sent the message, and the
	identity of the host that received the message (and is
	inserting this time stamp), and the date and time the
	message was received.  Relayed messages will have multiple
	time stamp lines.

	When the receiver-SMTP makes the "final delivery" of a
	message it inserts at the beginning of the mail data a
	return path line.  The return path line preserves the
	information in the <reverse-path> from the MAIL command.
	Here, final delivery means the message leaves the SMTP
	world.  Normally, this would mean it has been delivered to
	the destination user, but in some cases it may be further
	processed and transmitted by another mail system.

	   It is possible for the mailbox in the return path be
	   different from the actual sender's mailbox, for example,
	   if error responses are to be delivered a special error
	   handling mailbox rather than the message senders.

	The preceding two paragraphs imply that the final mail data
	will begin with a  return path line, followed by one or more
	time stamp lines.  These lines will be followed by the mail
	data header and body [2].  See Example 8.

	Special mention is needed of the response and further action
	required when the processing following the end of mail data
	indication is partially successful.  This could arise if
	after accepting several recipients and the mail data, the
	receiver-SMTP finds that the mail data can be successfully
	delivered to some of the recipients, but it cannot be to
	others (for example, due to mailbox space allocation
	problems).  In such a situation, the response to the DATA
	command must be an OK reply.  But, the receiver-SMTP must
	compose and send an "undeliverable mail" notification
	message to the originator of the message.  Either a single
	notification which lists all of the recipients that failed
	to get the message, or separate notification messages must
	be sent for each failed recipient (see Example 7).  All
	undeliverable mail notification messages are sent using the
	MAIL command (even if they result from processing a SEND,
	SOML, or SAML command).

DATA <CRLF>

I: 354 -> data -> S: 250
I: 354 Start mail input; end with <CRLF>.<CRLF>
S: 250 Requested mail action okay, completed
F: 552, 554, 451, 452
F: 451, 554
E: 500, 501, 503, 421
*/
type CmdDATA struct {
	connection net.Conn
	args       string
}

func NewCmdDATA(c net.Conn, args string) *CmdDATA {
	cmd := &CmdDATA{
		connection: c,
		args:       args,
	}
	return cmd
}

func (cmd *CmdDATA) RunCMD() {
	if len(cmd.args) != 0 {
		cmd.connection.Write([]byte("501 Syntax error in parameters or arguments\r\n"))
		cmd.connection.Close()
	}
	cmd.connection.Write([]byte("354 Start mail input; end with <CRLF>.<CRLF>\r\n"))
	r := textproto.NewReader(bufio.NewReader(cmd.connection))
	var data bytes.Buffer
	data.ReadFrom(r.DotReader())
	cmd.connection.Write([]byte("250 Requested mail action okay, completed\r\n"))
	fmt.Printf("Recieved data:%s", data.String())
}
