package account

import (
	"bk/src/static"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context)  {
	var data struct {
		Name string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数格式错误",
			Data: nil,
		})
		return
	}

	fmt.Println("Prams",data.Name, data.Password)
	ok , token ,err := LoginAccount(data.Name, data.Password)
	if err != nil {
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  error.Error(err),
			Data: nil,
		})
		return
	}
	if ok == true {
		fmt.Println(token)
		c.JSON(200, static.Response{
			Code: 200,
			Msg:  "登录成功",
			Data: token,
		})
		return
	}else{
		c.JSON(404, static.Response{
			Code: 404,
			Msg:  "未知类型错误请检查后端源代码",
			Data: nil,
		})
		return
	}
}

func Register(c *gin.Context)  {
	var data struct {
		Name string `json:"username"`
		Password string `json:"password"`
		Email string `json:"email"`
		Phone string `json:"phone"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数格式错误",
			Data: nil,
		})
		return
	}

	ok,err := RegisterAccount(data.Name,data.Phone,data.Password,data.Email)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, static.Response{
			Code: 10400,
			Msg:  "服务器内部错误",
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "登录成功",
		Data: ok,
	})
	return
}

