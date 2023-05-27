package browse

import (
	"bk/src/static"
	"github.com/gin-gonic/gin"
)

func Browse(c *gin.Context) {
	authorization := c.Query("authorization")

	username, calculus, err := GetUserInfo(authorization)
	if err != nil {
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  "服务器内部错误: "+ err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: struct {
			Username string `json:"username"`
			Calculus int64 `json:"calculus"`
		}{Username:username, Calculus:calculus},
	})
}
