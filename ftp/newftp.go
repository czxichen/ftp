package ftp

import (
	"fmt"
	"net"
)

type Ftp struct {
	con net.Conn
	ip  string
}

func NewFtp(ip string) (*Ftp, error) {
	buf := make([]byte, 1024)
	con, err := net.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}
	n, err := con.Read(buf)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(buf[:n]))
	return &Ftp{con, ip}, nil
}
