package rfc821

import (
	"bytes"
	"fmt"
	"go-smtp-rcv/internal"
	"go-smtp-rcv/internal/config"
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
	client internal.I_SMTP_CLIENT
}

func NewSMTP_rfc821(serv_conf config.Server) *SMTP_rfc821 {
	s := &SMTP_rfc821{}
	return s
}
func (s *SMTP_rfc821) SetClient(cl internal.I_SMTP_CLIENT) {
	s.client = cl
}
func (s *SMTP_rfc821) GetGreeating() string {
	return fmt.Sprintf("220 %s SMTP minimal rfc821", "!add domain!")
}
func (s *SMTP_rfc821) HandleCMD(mess internal.I_RawSMTPMessage) {
	//fmt.Printf("cmd: %s, args: %s\n", mess.GetSMTPCMD(), mess.GetCMDArgs())
	switch mess.GetSMTPCMD() {
	case "HELO":
		{
			NewCmdHELO(s.client, mess.GetCMDArgs()).RunCMD()
		}
	case "MAIL":
		{
			NewCmdMAIL(s.client, mess.GetCMDArgs()).RunCMD()
		}
	case "RCPT":
		{
			NewCmdRCPT(s.client, mess.GetCMDArgs()).RunCMD()
		}
	case "NOOP":
		{
			NewCmdNOOP(s.client).RunCMD()
		}
	case "DATA":
		{
			NewCmdDATA(s.client, mess.GetCMDArgs()).RunCMD()
		}
	case "QUIT":
		{
			NewCmdQUIT(s.client).RunCMD()
		}
	default:
		{
			s.client.GetSMTPConnection().WriteCMD("502 Command not implemented")
		}
	}

}
