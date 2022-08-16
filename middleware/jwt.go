package middleware

import (
	"gin-blog/utils"
	e "gin-blog/utils/errmsg"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JWT.SigningKey)

const (
	AuthorizationKey = "Authorization"
	BearerKey        = "Bearer"
	Null             = ""
	SpaceKey         = " "
	Num              = 2
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// SetToken 生成token
func SetToken(username string) (string, e.ResCode) {
	expireTime := utils.JWT.ExpiresTime
	SetClaims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime,
			Issuer:    utils.JWT.Issuer,
		},
	}
	reqClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	token, err := reqClaim.SignedString(JwtKey)
	if err != nil {
		return "", e.ERROR
	}
	return token, e.SUCCESS
}

// CheckToken 验证token
func CheckToken(data string) (*MyClaims, e.ResCode) {
	token, _ := jwt.ParseWithClaims(data, &MyClaims{}, func(data *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, _ := token.Claims.(*MyClaims); token.Valid {
		return key, e.SUCCESS
	} else {
		return nil, e.ERROR
	}
}

// JwtMiddleware 中间件
func JwtMiddleware() gin.HandlerFunc {
	code := e.SUCCESS
	return func(ctx *gin.Context) {
		tokenHeader := ctx.Request.Header.Get(AuthorizationKey)
		if tokenHeader == Null {
			code = e.ErrorTokenExist
			utils.ResponseErrorWithMsg(ctx, code, code.GetMsg())
			ctx.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, SpaceKey, Num)
		if len(checkToken) != 2 && checkToken[0] != BearerKey {
			code = e.ErrorTokenTypeWrong
			utils.ResponseErrorWithMsg(ctx, code, code.GetMsg())
			ctx.Abort()
			return
		}
		key, checkCode := CheckToken(checkToken[1])
		if checkCode == e.ERROR {
			code = e.ErrorTokenWrong
			utils.ResponseErrorWithMsg(ctx, code, code.GetMsg())
			ctx.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = e.ErrorTokenRuntime
			utils.ResponseErrorWithMsg(ctx, code, code.GetMsg())
			ctx.Abort()
			return
		}
		ctx.Set("username", key.Username)
		ctx.Next()
	}
}
