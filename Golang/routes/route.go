package routes

import (
	"github.com/gin-gonic/gin"
	"web/api"
)

func InitRouter(g *gin.Engine, api api.API) {
	g.POST("/api/login", api.Login)
	g.POST("/wechat", api.SendWechat)
	r := g.Group("/api")
	r.Use(api.Auth)
	r.GET("/zfs", api.ZFSInfo)
	r.GET("/disk", api.DiskInfo)
	r.GET("/crontab/:uid", api.GetCrontab)
	r.POST("/crontab/:uid", api.SetCrontab)
	r.GET("/wechat", api.GetWechat)
	r.POST("/wechat", api.SetWechat)
	r.GET("/wechat/test", api.TestWechat)
	r.GET("/reboot", api.Reboot)

}
