package utilies

import (
	"github.com/Ricky-fight/car-admin-server/core"
	"github.com/Ricky-fight/car-admin-server/global"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Bind(r interface{}, c *gin.Context) error {
	if err := c.ShouldBindJSON(r); err != nil {
		core.FailWithErr(core.BAD_REQUEST_ERROR, err, c)
		return err
	}
	return nil
}
func Valid(r interface{}, c *gin.Context) error {
	if err := global.VALIDATE.Struct(r); err != nil {
		core.FailWithValidation(err, c)
		return err
	}
	return nil
}

func BindAndValid(r interface{}, c *gin.Context) error {
	if err := Bind(r, c); err != nil {
		return err
	}
	if err := Valid(r, c); err != nil {
		return err
	}
	return nil
}

func GetIDFromParams(c *gin.Context) (id uint, err error) {
	// get id
	var idStr string
	// fmt.Printf("c.Params: %v\n", c.Params)
	if idStr = c.Param("id"); idStr == "0" || idStr == "" {
		core.FailWithBadRequest(c)
		return
	}
	if intNum, err := strconv.Atoi(idStr); err != nil {
		return 0, err
	} else {
		id = uint(intNum)
	}
	return
}
