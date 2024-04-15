package service

import (
	"os"
	"strings"
)

func (s *service) CheckPWD(pwd string) bool {
	f, err := os.ReadFile(s.exPath + "pwd")
	if err != nil {
		return true
	}
	data := strings.Trim(string(f), "\n")
	return pwd == data
}
