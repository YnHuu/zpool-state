package middleware

import (
	"github.com/kataras/jwt"
	"time"
)

var (
	secret []byte
	expire int64
)

type JwtToken struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

func InitJWT(accessSecret string, accessExpire int64) {
	secret = []byte(accessSecret)
	expire = accessExpire * 60
}

func GenerateToken() JwtToken {
	nowUnix := time.Now().Unix()
	token, _ := jwt.Sign(jwt.HS256, secret, jwt.Claims{
		Issuer:   "YnHuu",
		IssuedAt: nowUnix,
		Expiry:   nowUnix + expire,
	})
	return JwtToken{
		Token:  string(token),
		Expire: time.Now().Unix(),
	}
}

func VerifyToken(token string) (*jwt.VerifiedToken, error) {
	return jwt.Verify(jwt.HS256, secret, []byte(token))
}
