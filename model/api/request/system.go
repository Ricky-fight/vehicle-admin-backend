package request

type Login struct {
	Account  string `json:"account" validate:"required"`  // 用户名
	Password string `json:"password" validate:"required"` // 密码
}
