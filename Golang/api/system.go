package api

import "github.com/gin-gonic/gin"

func (a *api) Reboot(c *gin.Context) {
	a.resp(c, 200, nil)
	a.ser.Reboot()
}
