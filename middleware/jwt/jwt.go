package jwt

type Claims struct {
	Account   string `json:"sub"`
	NotBefore int64  `json:"nbf,omitempty"`
	ExpiresAt int64  `json:"exp,omitempty"`
	// IP   string // TODO 获取ip作为验证依据，防止重放攻击
	// jwt.StandardClaims
}

func (c Claims) Valid() error {
	return nil
}

// jwt简介：
// https://www.jianshu.com/p/576dbf44b2ae
// iss: jwt签发者
// sub: jwt所面向的用户
// aud: 接收jwt的一方
// exp: jwt的过期时间，这个过期时间必须要大于签发时间
// nbf: 定义在什么时间之前，该jwt都是不可用的.
// iat: jwt的签发时间
// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。

// // 读取配置
// var viperConfig = viper.GetStringMap()
// var TokenExpireDuration = viper.Get("jwt.TokenExpireDuration")
// var secret = []byte(viper.Get("jwt.TokenExpi"))

// func JwtAuth(c *gin.Context) gin.HandlerFunc {
// 	var claims Claims
// 	token := jwt.Parse(c.GetHeader())
// }
