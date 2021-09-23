package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

// PageResult 查询结果列表，返回时作为data嵌入Response
type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type EmptyData map[string]interface{}

const (
	ERROR   = -1
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "success", c)
}
func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, EmptyData{}, message, c)
}
func Ok(c *gin.Context) {
	Result(SUCCESS, EmptyData{}, "success", c)
}
func Fail(c *gin.Context) {
	Result(ERROR, EmptyData{}, "fail", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, EmptyData{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}
