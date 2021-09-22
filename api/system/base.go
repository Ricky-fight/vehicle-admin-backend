package system

import (
	"fmt"
	"regexp"
	"time"

	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/global"
	jwtMiddleware "github.com/Ricky-fight/car-admin-server/middleware/jwt"
	"github.com/Ricky-fight/car-admin-server/model/database"
	"github.com/Ricky-fight/car-admin-server/service/system"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginResp struct {
	Account   string
	Token     string
	ExpiresAt int64
}
type LoginReq struct {
	Account  string `json:"account" validate:"required"`  // 用户名
	Password string `json:"password" validate:"required"` // 密码
}

func init() {
	global.VALIDATE.RegisterValidation("password", validatePassword) // 注册自定义函数，前一个参数是struct里tag自定义，后一个参数是自定义的函数
}

// 用户登录
func Login(c *gin.Context) {
	// bind data
	loginParam := LoginReq{}

	if err := c.ShouldBindJSON(&loginParam); err != nil {
		core.FailWithErr(core.BAD_REQUEST_ERROR, err, c)
		return
	}
	// validate
	if err := global.VALIDATE.Struct(&loginParam); err != nil {
		core.FailWithValidation(err, c)
		return
	}
	// prepare data for service
	u := &database.User{
		Account:  loginParam.Account,
		Password: loginParam.Password,
	}
	// deal with login
	if err := system.Login(u); err != nil {
		if loginResponse, err := SigningJWTToken(loginParam); err != nil {
			core.FailWithErr(core.AUTH_ERROR, err, c)
		} else {
			core.OkWithDetailed(loginResponse, "login success", c)
		}
	} else {
		core.FailWithErr(core.AUTH_ERROR, fmt.Errorf("wrong account or password"), c)
	}
}

// 生成jwt令牌
func SigningJWTToken(user LoginReq) (LoginResp, error) {
	now := time.Now()
	expireDuration := global.CONFIG.Jwt.TokenExpireDuration
	claims := jwtMiddleware.Claims{
		Account: user.Account,
		StandardClaims: jwt.StandardClaims{
			NotBefore: now.Add(time.Second * -1).Unix(),
			ExpiresAt: now.Add(expireDuration).Unix(),
		},
	}
	// fmt.Printf("claims: %+v\n", claims)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// fmt.Printf("token: %v\n", token)
	secret := []byte(global.CONFIG.Jwt.Secret)
	// fmt.Printf("secret: %v\n", secret)

	if xToken, err := token.SignedString(secret); err != nil {
		return LoginResp{}, err
	} else {
		l := LoginResp{
			Account:   user.Account,
			Token:     xToken,
			ExpiresAt: claims.ExpiresAt * 1000,
		}
		return l, nil
	}
}

type RegisterReq struct {
	Account         string `json:"account" validate:"required,printascii,min=3,max=20"`  // 用户名
	Password        string `json:"password" validate:"required,printascii,password"`     // 密码
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"` // 二次确认密码
}

// 密码校验，用于validator字段tag
func validatePassword(passwordField validator.FieldLevel) bool {
	// 密码只含有字母数字下划线，长度6-20
	reg := `^(.*[\w]){6,20}$`
	password := passwordField.Field().String()
	rst, _ := regexp.MatchString(reg, password)
	return rst
}

// 用户注册
func Register(c *gin.Context) {
	// bind data
	r := RegisterReq{}
	if err := c.ShouldBindJSON(&r); err != nil {
		core.FailWithErr(core.BAD_REQUEST_ERROR, err, c)
		return
	}
	// valid
	if err := global.VALIDATE.Struct(&r); err != nil {
		core.FailWithValidation(err, c)
		return
	}
	// prepare data for service
	user := database.User{
		Account:  r.Account,
		Password: r.Password,
	}
	if err := system.Register(&user); err != nil {
		core.FailWithErr(core.ERROR, err, c)
		return
	}
}
