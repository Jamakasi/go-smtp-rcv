package rfc821

import "go-smtp-rcv/internal"

/*
S: 221 <domain> Service closing transmission channel
E: 500
*/
type CmdQUIT struct {
	client internal.I_SMTP_CLIENT
}

func NewCmdQUIT(c internal.I_SMTP_CLIENT) *CmdQUIT {
	cmd := &CmdQUIT{
		client: c,
	}
	return cmd
}

func (cmd *CmdQUIT) RunCMD() {
	cmd.client.GetSMTPConnection().WriteCMD("221 <domain> Service closing transmission channel")
}
