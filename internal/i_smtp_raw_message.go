package internal

type I_RawSMTPMessage interface {
	GetSMTPCMD() string
	GetCMDArgs() string
	GetErrCode() uint
	GetErr() error
}
