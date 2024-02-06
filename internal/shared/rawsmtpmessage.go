package shared

type RawSMTPMessage struct {
	smtp_cmd string
	cmd_args string
	errcode  uint
	err      error
}

const (
	R_OK_WITH_ARGS    = iota
	R_OK_WITHOUT_ARGS = iota
	R_E_NOT_CMD       = iota
	R_E_READ_ERROR    = iota
	W_OK              = iota
)

func (m *RawSMTPMessage) GetSMTPCMD() string {
	return m.smtp_cmd
}
func (m *RawSMTPMessage) GetCMDArgs() string {
	return m.cmd_args
}
func (m *RawSMTPMessage) GetErrCode() uint {
	return m.errcode
}
func (m *RawSMTPMessage) GetErr() error {
	return m.err
}
