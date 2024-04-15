package api

import (
	"github.com/gin-gonic/gin"
	"io"
)

func (a *api) GetCrontab(c *gin.Context) {
	uid := c.Param("uid")
	if uid != "" {
		data := a.ser.GetCrontab(uid)
		a.resp(c, 200, data)
		return
	}
	a.resp(c, 20010, nil)
}

func (a *api) SetCrontab(c *gin.Context) {
	uid := c.Param("uid")
	if uid == "" {
		a.resp(c, 20010, nil)
		return
	}

	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		a.resp(c, 20011, nil)
		return
	}
	if a.ser.SetCrontab(uid, b) {
		a.resp(c, 200, "保存成功")
		return
	}
	a.resp(c, 20012, nil)
}
