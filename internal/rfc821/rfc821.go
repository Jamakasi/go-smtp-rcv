package rfc821

import (
	"bytes"
	"fmt"
	"go-smtp-rcv/internal/config"
	"net"
)

const (
	SIZE_USER   = 64  //The maximum total length of a user name is 64 characters.
	SIZE_DOMAIN = 64  //The maximum total length of a domain name or number is 64 characters.
	SIZE_RFPATH = 256 /*The maximum total length of a reverse-path or
	forward-path is 256 characters (including the punctuation and element separators).*/
	SIZE_CMDLINE   = 512  //The maximum total length of a command line including the command word and the <CRLF> is 512 characters.
	SIZE_REPLYLINE = 512  //The maximum total length of a reply line including the reply code and the <CRLF> is 512 characters.
	SIZE_TEXTLINE  = 1000 /*The maximum total length of a text line including the
	<CRLF> is 1000 characters (but not counting the leading
	dot duplicated for transparency).*/
	SIZE_RCPTBUF = 100 //The maximum total number of recipients that must be buffered is 100 recipients.
)

type command []byte

// 4.5.1.  MINIMUM IMPLEMENTATION
var (
	cmdRSET command = []byte("RSET")
	cmdDATA command = []byte("DATA")
)

func (c command) match(in []byte) bool {
	return bytes.Index(in, []byte(c)) == 0
}

type SMTP_rfc821 struct {
	connection    net.Conn
	server_config config.Server
}

func NewSMTP_rfc821(con net.Conn, serv_conf config.Server) *SMTP_rfc821 {
	s := &SMTP_rfc821{
		connection:    con,
		server_config: serv_conf,
	}
	return s
}
func (s *SMTP_rfc821) GetConn() net.Conn {
	return s.connection
}
func (s *SMTP_rfc821) GetGreeating() string {
	return fmt.Sprintf("220 %s SMTP minimal rfc821\r\n", s.server_config.S_domain)
}
func (s *SMTP_rfc821) HandleCMD(cmd string, vars string) {

	switch cmd {
	case "HELO":
		{
			NewCmdHELO(s.connection, vars).RunCMD()
		}
	case "MAIL":
		{
			NewCmdMAIL(s.connection, vars).RunCMD()
		}
	case "RCPT":
		{
			NewCmdRCPT(s.connection, vars).RunCMD()
		}
	case "NOOP":
		{
			NewCmdNOOP(s.connection).RunCMD()
		}
	case "DATA":
		{
			NewCmdDATA(s.connection, vars).RunCMD()
		}
	case "QUIT":
		{
			NewCmdQUIT(s.connection).RunCMD()
		}
	default:
		{
			s.connection.Write([]byte("502 Command not implemented\r\n"))
		}
	}

}
