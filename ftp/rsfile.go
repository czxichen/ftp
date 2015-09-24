package ftp

import (
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func (self *Ftp) PutPasv(Pathname string) error {
	con, err := self.connection("STOR", Pathname)
	if err != nil {
		return err
	}
	File, err := os.Open(Pathname)
	if err != nil {
		con.Close()
		return err
	}
	io.Copy(con, File)
	File.Close()
	con.Close()
	buf := make([]byte, 1024)
	_, err = self.con.Read(buf)
	if err != nil {
		return err
	}
	return nil
}
func (self *Ftp) GetFile(Pathname string) error {
	con, err := self.connection("RETR", Pathname)
	if err != nil {
		return err
	}
	File, err := os.Create(Pathname)
	if err != nil {
		con.Close()
		return err
	}
	io.Copy(File, con)
	File.Close()
	con.Close()
	buf := make([]byte, 1024)
	_, err = self.con.Read(buf)
	if err != nil {
		return err
	}
	return nil
}
func (self *Ftp) connection(status, Pathname string) (net.Conn, error) {
	buf := make([]byte, 1024)
	self.con.Write([]byte("PASV \r\n"))
	n, err := self.con.Read(buf)
	if err != nil {
		return nil, err
	}
	if s := string(buf[:n]); !strings.Contains(s, "227 Entering Passive Mode") {
		return nil, errors.New(s)
	}
	port := getport(buf[27 : n-3])
	con, err := net.Dial("tcp", fmt.Sprintf("%s:%d", strings.Split(self.ip, ":")[0], port))
	if err != nil {
		return nil, err
	}
	self.con.Write([]byte(fmt.Sprintf("%s %s\r\n", status, Pathname)))
	n, err = self.con.Read(buf)
	if err != nil {
		con.Close()
		return nil, err
	}
	if !strings.Contains(string(buf[:n]), "150 Opening data channel") {
		con.Close()
		return nil, errors.New("create data link error.")
	}
	return con, nil
}
