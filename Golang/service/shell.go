package service

import (
	"os/exec"
	"strings"
)

func shell(cmd string) string {
	sh := exec.Command("sh", "-c", cmd)
	bytes, err := sh.Output()
	if err != nil {
		return ""
	}
	return strings.Trim(string(bytes), "\n")
}
