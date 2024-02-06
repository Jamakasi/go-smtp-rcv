package shared

import (
	"go-smtp-rcv/internal"
	"go-smtp-rcv/internal/config"
	"go-smtp-rcv/internal/rfc821"

	"net"
)

type SMTPClient struct {
	id           int
	smtpcon      internal.I_SMTP_CONNECTION
	smtp_rfc_ins internal.I_SMTP_SPEC
	config       config.Server
}

func NewSMTPClient(clid int, connection net.Conn, conf config.Server) *SMTPClient {
	smtpCon := NewSMTPConnection(connection)
	// определять из конфига
	stype := rfc821.NewSMTP_rfc821(conf)

	c := &SMTPClient{id: clid,
		smtpcon:      smtpCon,
		smtp_rfc_ins: stype,
		config:       conf}
	stype.SetClient(c)
	return c
}

func (cl *SMTPClient) GetID() int {
	return cl.id
}
func (cl *SMTPClient) GetSMTPConnection() internal.I_SMTP_CONNECTION {
	return cl.smtpcon
}

func (cl *SMTPClient) Handle() {
	cl.smtpcon.WriteCMD(cl.smtp_rfc_ins.GetGreeating())
	for {
		smtp_mess := cl.smtpcon.ReadCMD()
		switch smtp_mess.GetErrCode() {
		case R_E_READ_ERROR:
			{
				//log log log
				return
			}
		default:
			{
				cl.smtp_rfc_ins.HandleCMD(smtp_mess)
			}
		}

	}
}
