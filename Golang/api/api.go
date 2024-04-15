package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"web/service"
)

type (
	API interface {
		Auth(*gin.Context)
		Login(*gin.Context)
		ZFSInfo(*gin.Context)
		DiskInfo(*gin.Context)
		GetCrontab(*gin.Context)
		SetCrontab(*gin.Context)
		GetWechat(*gin.Context)
		SetWechat(*gin.Context)
		TestWechat(*gin.Context)
		SendWechat(*gin.Context)
		Reboot(*gin.Context)
	}
	api struct {
		ser service.Service
	}
)

func InitApi(ser service.Service) API {
	return &api{ser}
}

type apiResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data,omitempty"`
	Msg  string      `json:"msg,omitempty"`
}

var errCode = map[int]string{
	10000: "密码错误",
	10010: "json解析错误",

	20010: "定时任务类型错误",
	20011: "定时任务流处理失败",
	20012: "定时任务保存失败",

	30010: "微信推送流处理失败",
	30011: "微信推送流保存失败",
	30020: "微信推送发送失败",
	30021: "不支持对外IP地址",
}

func (*api) resp(c *gin.Context, code int, data any) {
	c.JSON(http.StatusOK, &apiResp{
		Code: code,
		Data: data,
		Msg:  errCode[code],
	})
}
