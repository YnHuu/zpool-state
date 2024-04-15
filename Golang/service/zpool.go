package service

import (
	"strings"
)

type Zpool struct {
	Pool  string `json:"pool,omitempty"`
	State string `json:"state,omitempty"`
	Raid  string `json:"raid,omitempty"`
	Size  string `json:"size,omitempty"`
	Alloc string `json:"alloc,omitempty"`
	Free  string `json:"free,omitempty"`
	Cap   string `json:"cap,omitempty"`
	Items []item `json:"items,omitempty"`
}

type item struct {
	Name  string `json:"name,omitempty"`
	State string `json:"state,omitempty"`
	Read  string `json:"read,omitempty"`
	Write string `json:"write,omitempty"`
	Cksum string `json:"cksum,omitempty"`
}

func pool() []string {
	data := shell(`zpool list -Ho name`)
	zpools := strings.Split(data, "\n")
	return zpools
}

func poolStatus(p string) Zpool {
	data := shell(`zpool list -Ho name,size,alloc,free,cap,health ` + p)
	list := strings.Split(data, "\t")
	if len(list) != 6 {
		return Zpool{}
	}

	data = shell(`zpool status ` + p + `|awk '/config:/,/errors:/'|awk 'NR>4 {print $1,tolower($2),$3,$4,$5}'|head -n -2`)
	status := strings.Split(data, "\n")

	var items []item
	raid := ""
	for i, dev := range status {
		devStatus := strings.Split(dev, " ")
		if len(devStatus) != 5 {
			continue
		}
		if i == 0 {
			raid = devStatus[0]
			continue
		}
		items = append(items, item{
			Name:  devStatus[0],
			State: devStatus[1],
			Read:  devStatus[2],
			Write: devStatus[3],
			Cksum: devStatus[4],
		})
	}
	zpool := Zpool{
		Pool:  p,
		State: list[5],
		Raid:  raid,
		Size:  list[1],
		Alloc: list[2],
		Free:  list[3],
		Cap:   strings.Replace(list[4], "%", "", -1),
		Items: items,
	}

	return zpool
}

func (s *service) GetPoolInfo() any {
	pools := pool()
	var zpools []Zpool
	for _, pool := range pools {
		zpools = append(zpools, poolStatus(pool))
	}
	//log.Printf("%+v\n", zpools)
	return zpools
}
