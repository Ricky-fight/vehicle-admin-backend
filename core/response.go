package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var EmptyData = map[string]interface{}{}

const (
	// 状态码
	SUCCESS = 000 + iota
	ERROR   = 100 + iota
	VALIDATE_ERROR
	BAD_REQUEST_ERROR
	AUTH_ERROR
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, EmptyData, "success", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, EmptyData, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, EmptyData, "fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, EmptyData, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithValidation(data interface{}, c *gin.Context) {
	Result(VALIDATE_ERROR, data, "fail to validate data", c)
}

func FailWithBadRequest(c *gin.Context) {
	Result(BAD_REQUEST_ERROR, EmptyData, "bad request", c)
}
func FailWithErr(code int, err error, c *gin.Context) {
	Result(code, EmptyData, err.Error(), c)
}
