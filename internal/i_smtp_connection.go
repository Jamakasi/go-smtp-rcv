package internal

type I_SMTP_CONNECTION interface {
	ReadCMD() I_RawSMTPMessage
	GetSMTPConnection() I_SMTP_CONNECTION
	WriteCMD(msg string) (n int, err error)
	ReadRawCRLF() (string, error)
	ReadRawCRLFDotCRLF() (string, error)
}
