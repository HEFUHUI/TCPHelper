package base

import (
	"bytes"
	"net"
)

type FileReciver struct {
	host      string
	timeout   int64
	conn      *net.Conn
	file_type string
}

func (f *FileReciver) writeFile(b bytes.Buffer) {
	// TODO
}
