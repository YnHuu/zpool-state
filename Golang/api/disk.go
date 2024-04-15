package api

import "github.com/gin-gonic/gin"

func (a *api) DiskInfo(c *gin.Context) {
	disks := a.ser.GetDiskInfo()
	a.resp(c, 200, disks)
}
