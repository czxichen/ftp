package ftp

import (
	"errors"
	"fmt"
	"strings"
)

func (self *Ftp) Login(user, pass string) error {
	buf := make([]byte, 1024)
	self.con.Write([]byte(fmt.Sprintf("USER %s\r\n", user)))
	self.con.Read(buf)
	self.con.Write([]byte(fmt.Sprintf("PASS %s\r\n", pass)))
	n, err := self.con.Read(buf)
	if err != nil {
		return err
	}
	if !strings.Contains(string(buf[:n]), "230 Logged on") {
		return errors.New(strings.TrimSpace(string(buf)))
	}
	return nil
}
