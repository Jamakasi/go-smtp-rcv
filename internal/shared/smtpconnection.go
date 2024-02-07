package shared

import (
	"fmt"
	"go-smtp-rcv/internal"
	"net"
	"strings"
)

type SMTPConnection struct {
	connection  net.Conn
	recievedBuf []byte
}

func NewSMTPConnection(con net.Conn) *SMTPConnection {
	i := &SMTPConnection{connection: con}
	return i
}
func (mw *SMTPConnection) GetSMTPConnection() internal.I_SMTP_CONNECTION {
	return mw
}

/*
*
Read line ends with CRLF (\r\n)
Return CMD in uppercase, AGRS in raw string, ERRCode in uint
ARGS may be empty
ERRCode:

	OK_WITHOUT_ARGS
	OK_WITH_ARGS
	NOT CMD. Line shorten than 4 chars
	Read error
*/
func (mw *SMTPConnection) ReadCMD() internal.I_RawSMTPMessage {
	raw, err := mw.ReadRawCRLF() //data.ReadFrom(r.DotReader())

	if err != nil {
		return &RawSMTPMessage{"", "", R_E_READ_ERROR, err}
	}
	if len(raw) < 4 {
		return &RawSMTPMessage{raw, "", R_E_NOT_CMD, fmt.Errorf("line\"%s\" < 4 chars", raw)}
	}
	msg := strings.ToUpper(raw[:4])
	if len(raw) > 5 {
		return &RawSMTPMessage{msg, raw[5:], R_OK_WITH_ARGS, nil}
	}
	return &RawSMTPMessage{msg, "", R_OK_WITHOUT_ARGS, nil}
}

/*
func (mw *MsgWorker) ReadCRLF() (msg string, args string, errcode uint, err error) {
	reader := textproto.NewReader(bufio.NewReader(mw.connection))
	raw, err := reader.ReadLine() //data.ReadFrom(r.DotReader())
	if err != nil {
		return "", "", R_E_READ_ERROR, err
	}
	if len(raw) < 4 {
		return raw, "", R_E_NOT_CMD, fmt.Errorf("line\"%s\" < 4 chars", raw)
	}
	msg = strings.ToUpper(raw[:4])
	if len(raw) > 5 {
		return msg, raw[5:], R_OK_WITH_ARGS, nil
	}
	return msg, "", R_OK_WITHOUT_ARGS, nil
}
*/

func (mw *SMTPConnection) WriteCMD(msg string) (n int, err error) {
	if !strings.HasSuffix(msg, "\r\n") {
		msg += "\r\n"
	}
	for len(msg) > 0 {
		n, err = mw.connection.Write([]byte(msg))
		if err != nil {
			return n, err
		}
		msg = msg[n:]
		//fmt.Printf("write: %s err:%s\n", msg, err)
	}

	return 0, nil
}

func (mw *SMTPConnection) ReadRawCRLF() (string, error) {
	var buf []byte
	for {
		b := make([]byte, 1024)
		n, err := mw.connection.Read(b)
		if err != nil {
			return "", err
		}

		buf = append(buf, b[:n]...)
		for i, b := range buf {
			// If end of line
			if b == '\n' && i > 0 && buf[i-1] == '\r' {
				// i-1 because drop the CRLF, no one cares after this
				line := string(buf[:i-1])
				//buf = buf[i+1:]
				mw.recievedBuf = append(mw.recievedBuf, buf...)
				return line, nil
			}
		}
	}
}
func (mw *SMTPConnection) ReadRawCRLFDotCRLF() (string, error) {
	var buf []byte
	for {
		b := make([]byte, 1024)
		n, err := mw.connection.Read(b)
		if err != nil {
			return "", err
		}

		buf = append(buf, b[:n]...)
		for i, b := range buf {
			// If end of line
			if i > 4 && b == '\n' && buf[i-1] == '\r' && buf[i-2] == '.' && buf[i-3] == '\n' && buf[i-4] == '\r' {
				// i-4 because drop the CRLF, no one cares after this
				line := string(buf[:i-4])
				//buf = buf[i+4:]
				mw.recievedBuf = append(mw.recievedBuf, buf...)
				return line, nil
			}
		}
	}
}
