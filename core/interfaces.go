package core

import "github.com/gin-gonic/gin"

type ApiRouter interface {
	Register(*gin.RouterGroup)
}
