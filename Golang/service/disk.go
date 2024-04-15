package service

import (
	"fmt"
	"strconv"
	"strings"
)

type Disk struct {
	Index           int    `json:"index,omitempty"`
	Use             string `json:"use,omitempty"`
	ModelName       string `json:"model_name,omitempty"`
	UserCapacity    string `json:"user_capacity,omitempty"`
	SmartStatus     string `json:"smart_status,omitempty"`
	Dev             string `json:"dev,omitempty"`
	Temperature     string `json:"temperature,omitempty"`
	SerialNumber    string `json:"serial_number,omitempty"`
	FirmwareVersion string `json:"firmware_version,omitempty"`
	PowerOnTime     string `json:"power_on_time,omitempty"`
	FormFactor      string `json:"form_factor,omitempty"`
}

func disk(dev string) []string {
	data := shell(`smartctl -a /dev/` + dev + ` -j|jq -r '.model_name,.smart_status.passed,.temperature.current,.serial_number,.firmware_version,.user_capacity.bytes,.power_on_time.hours,.form_factor.name'`)
	info := strings.Split(data, "\n")
	return info
}

func diskInfo(index int, p string) []Disk {
	var disks []Disk
	data := shell(`zpool list -Hv ` + p + `|awk 'NR>2 {print $1}'`)
	for i, d := range strings.Split(data, "\n") {
		info := disk(d)
		if len(info) != 8 {
			continue
		}
		disks = append(disks, Disk{
			Index:           i + index,
			Use:             p,
			ModelName:       info[0],
			UserCapacity:    b2e(info[5]),
			SmartStatus:     info[1],
			Dev:             d,
			Temperature:     info[2],
			SerialNumber:    info[3],
			FirmwareVersion: info[4],
			PowerOnTime:     info[6],
			FormFactor:      info[7],
		})
	}
	return disks
}

func b2e(str string) string {
	dmap := map[int]string{1: "KB", 2: "MB", 3: "GB", 4: "TB", 5: "PB", 6: "EB"}
	n, _ := strconv.Atoi(str)
	c := 0
	for n > 1000 {
		n /= 1000
		c++
	}
	data := fmt.Sprintf("%.1f %v", float64(n), dmap[c])
	return data
}

func (s *service) GetDiskInfo() any {
	pools := pool()
	var disks []Disk
	for _, pool := range pools {
		index := len(disks) + 1
		disks = append(disks, diskInfo(index, pool)...)
	}
	//log.Printf("%+v\n", disks)
	return disks
}
