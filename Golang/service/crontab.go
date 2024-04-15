package service

import (
	"log"
	"os"
)

func (s *service) GetCrontab(uid string) any {
	//15min    daily    hourly   monthly  weekly
	f, err := os.ReadFile("/etc/periodic/" + uid + "/cron")
	if err != nil {
		log.Println(err)
		return nil
	}
	return string(f)
}

func (s *service) SetCrontab(uid string, data []byte) bool {
	if len(data) == 0 {
		_ = os.Remove("/etc/periodic/" + uid + "/cron")
		return true
	}
	return os.WriteFile("/etc/periodic/"+uid+"/cron", data, 0755) == nil
}
