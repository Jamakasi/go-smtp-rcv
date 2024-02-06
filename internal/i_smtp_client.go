package internal

type I_SMTP_CLIENT interface {
	GetID() int
	GetSMTPConnection() I_SMTP_CONNECTION
}
