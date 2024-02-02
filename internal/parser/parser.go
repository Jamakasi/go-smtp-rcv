package parser

import (
	"bufio"
	"fmt"
	"go-smtp-rcv/internal"
	"net/textproto"
	"strings"
)

type Parser struct {
	Ver internal.IF_SMTP_STD
}

func (p *Parser) Parse() {
	r := textproto.NewReader(bufio.NewReader(p.Ver.GetConn()))
	p.Ver.GetConn().Write([]byte(p.Ver.GetGreeating()))
	for {
		//var data bytes.Buffer
		raw, err := r.ReadLine() //data.ReadFrom(r.DotReader())
		if err != nil {
			fmt.Printf("Read error: %s\n", err)
			break
		}
		//raw := data.String()
		if len(raw) < 4 {
			fmt.Printf("Unknown data: %s\n", raw)
			continue
		}
		cmd := strings.ToUpper(raw[:4])
		var args string
		if len(raw) > 5 {
			args = raw[5:]
		} else {
			args = ""
		}
		fmt.Printf("Parsed cmd: \"%s\" args: \"%s\"\n", cmd, args)
		p.Ver.HandleCMD(cmd, args)
	}
}
