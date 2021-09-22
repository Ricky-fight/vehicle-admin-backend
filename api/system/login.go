package system

import (
	"fmt"
	"time"

	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/global"
	jwtMiddleware "github.com/Ricky-fight/car-admin-server/middleware/jwt"
	"github.com/Ricky-fight/car-admin-server/model/api/request"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"github.com/Ricky-fight/car-admin-server/service/system"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type LoginResponse struct {
	Account   string
	Token     string
	ExpiresAt int64
}

// 用户登录
func Login(c *gin.Context) {
	// bind data
	loginParam := request.Login{}

	if err := c.ShouldBindJSON(&loginParam); err != nil {
		core.FailWithMessage(err.Error(), c)
		return
	}
	// validate

	if err := global.VALIDATE.Struct(loginParam); err != nil {
		core.FailWithValidation(loginParam, c)
		return
	}
	// prepare data for service
	u := &database.User{
		Account:  loginParam.Account,
		Password: loginParam.Password,
	}
	// deal with login
	if rst := system.Login(u); rst {
		now := time.Now()
		expireDuration := global.CONFIG.Jwt.TokenExpireDuration
		claims := jwtMiddleware.Claims{
			Account:   loginParam.Account,
			NotBefore: now.Add(time.Second * -1).Unix(),
			ExpiresAt: now.Add(expireDuration).Unix(),
		}
		fmt.Printf("claims: %+v\n", claims)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		fmt.Printf("token: %v\n", token)
		secret := []byte(global.CONFIG.Jwt.Secret)
		fmt.Printf("secret: %v\n", secret)
		xToken, err := token.SignedString(secret)
		if err != nil {
			core.FailWithMessage("生成JWT签名时发生错误: "+err.Error(), c)
		}
		l := LoginResponse{
			Account:   loginParam.Account,
			Token:     xToken,
			ExpiresAt: claims.ExpiresAt * 1000,
		}
		core.OkWithData(l, c)
		return
	} else {
		core.FailWithAuth(loginParam, c)
		return
	}

}
