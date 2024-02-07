package internal

type I_LOGGER interface {
	LError(err error)
	LInfo(str string)
	LDebug(str string)
}
