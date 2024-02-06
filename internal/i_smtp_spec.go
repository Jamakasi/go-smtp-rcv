package internal

type I_SMTP_SPEC interface {
	GetGreeating() string
	HandleCMD(mess I_RawSMTPMessage)
	SetClient(cl I_SMTP_CLIENT)
}
