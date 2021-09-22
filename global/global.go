package global

import (
	"github.com/Ricky-fight/car-admin-server/model/config"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var (
	CONFIG   *config.Config
	DB       *gorm.DB
	VALIDATE *validator.Validate = validator.New()
	ROUTER   *gin.Engine
)

const ()
