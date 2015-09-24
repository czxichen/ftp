package ftp

import (
	"strconv"
	"strings"
)

func getport(by []byte) int {
	s := string(by)
	list := strings.Split(s, ",")
	n1, err := strconv.Atoi(list[len(list)-2])
	if err != nil {
		return 0
	}
	n2, err := strconv.Atoi(list[len(list)-1])
	if err != nil {
		return 0
	}
	return n1*256 + n2
}
