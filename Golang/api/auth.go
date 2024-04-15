package api

import (
	"github.com/gin-gonic/gin"
	"web/middleware"
)

func (a *api) Auth(c *gin.Context) {
	token := c.GetHeader("Token")

	_, err := middleware.VerifyToken(token)
	if err != nil {
		a.resp(c, 401, err.Error())
		c.Abort()
		return
	}
	c.Next()
}
