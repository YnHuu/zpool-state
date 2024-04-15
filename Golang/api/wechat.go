package api

import (
	"github.com/gin-gonic/gin"
	"web/models"
)

func (a *api) GetWechat(c *gin.Context) {
	data := a.ser.GetWechat()
	a.resp(c, 200, data)
}

func (a *api) SetWechat(c *gin.Context) {
	data := models.Wechat{}
	if err := c.BindJSON(&data); err != nil {
		a.resp(c, 10010, err.Error())
		return
	}
	if a.ser.SetWechat(data) {
		a.resp(c, 200, "保存成功")
		return
	}
	a.resp(c, 30011, nil)
}

func (a *api) TestWechat(c *gin.Context) {
	if a.ser.TestWechat() {
		a.resp(c, 200, "发送成功")
	} else {
		a.resp(c, 30020, nil)
	}
}

func (a *api) SendWechat(c *gin.Context) {
	cip := c.ClientIP()
	if cip == "127.0.0.1" || cip == "::1" {
		task := c.PostForm("task")
		status := c.PostForm("status")
		if a.ser.SendWechat(task, status) {
			c.String(200, "ok")
		} else {
			c.String(200, "err")
		}
		return
	}
	a.resp(c, 30021, nil)
}
