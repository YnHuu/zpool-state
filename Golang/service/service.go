package service

import (
	"os"
	"path/filepath"
	"web/models"
)

type (
	Service interface {
		CheckPWD(string) bool
		Reboot()
		GetPoolInfo() any
		GetDiskInfo() any
		GetCrontab(string) any
		SetCrontab(string, []byte) bool
		GetWechat() models.Wechat
		SetWechat(models.Wechat) bool
		TestWechat() bool
		SendWechat(string, string) bool
	}
	service struct {
		exPath string
	}
)

func InitService() Service {
	ex, _ := os.Executable()
	exPath := filepath.Dir(ex) + "/"
	return &service{
		exPath: exPath,
	}
}
