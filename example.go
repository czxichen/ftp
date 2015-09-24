package main

import (
	"fmt"
	"ftp"
)

func main() {
	Ftp, err := ftp.NewFtp("127.0.0.1:21")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Ftp.Login("root", "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = Ftp.GetFile("config.cfg")
	if err != nil {
		fmt.Println(err)
	}
}
