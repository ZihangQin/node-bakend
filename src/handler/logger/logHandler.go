package logger

import (
	"bk/src/static"
	"bk/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoggerBody struct {
	OpName string `json:"opName"`
	UserName string  `json:"UserName"`
	LocaleDate string `json:"LocaleDate"`
	LocaleTime string `json:"LocaleTime"`
	OpHash string `json:"opHash"`
}
func Loggers(c *gin.Context)  {
	var log LoggerBody
	if err := c.Bind(&log); err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误：" + err.Error(),
			Data: nil,
		})
		return
	}
	fmt.Println(log)

	err := SetLogger(log.LocaleDate,log.LocaleTime,log.OpName,log.UserName,log.OpHash)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10100,
		Msg:  "success",
		Data: nil,
	})
	return
}

func GetLoggers(c *gin.Context)  {
	page := c.Query("page")
	pageInt, err := utils.StringToInt(page)
	if err != nil {
		c.JSON(400,static.Response{
			Code: 1400,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	loggerList,totlePages,err := GetLogger(pageInt)
	if err !=nil {
		c.JSON(500,static.Response{
			Code: 10500,
			Msg:  "获取日志列表失败，请重新尝试",
			Data: nil,
		})
		return
	}
	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: struct {
			Loggers interface{} `json:"loggers"`
			TotlePage int `json:"totle_page"`
		}{Loggers:loggerList,TotlePage:totlePages},
	})
}