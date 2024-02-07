package logstd

import (
	"log"
	"os"
)

type LogSTD struct {
	stdI *log.Logger
	stdE *log.Logger
}

func NewLogSTD() *LogSTD {
	li := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	le := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	return &LogSTD{stdI: li, stdE: le}
}

func (l *LogSTD) LInfo(str string) {
	l.stdI.Println(str)
}

func (l *LogSTD) LDebug(str string) {
	l.stdI.Println(str)
}

func (l *LogSTD) LError(err error) {
	l.stdE.Println(err)
}
