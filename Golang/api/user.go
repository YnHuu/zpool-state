package api

import (
	"github.com/gin-gonic/gin"
	"web/middleware"
)

type Claims struct {
	Password string `json:"password,option"`
}

func (a *api) Login(c *gin.Context) {
	claims := Claims{}
	if err := c.BindJSON(&claims); err != nil {
		a.resp(c, 10010, err.Error())
		return
	}
	if !a.ser.CheckPWD(claims.Password) {
		a.resp(c, 10000, nil)
		return
	}

	jwtToken := middleware.GenerateToken()
	a.resp(c, 200, jwtToken)
	return

}
