package api

import "github.com/gin-gonic/gin"

func (a *api) ZFSInfo(c *gin.Context) {
	zpools := a.ser.GetPoolInfo()
	a.resp(c, 200, zpools)
}
