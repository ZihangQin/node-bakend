package TestList

import (
	"bk/src/static"
	"bk/src/utils"
	"github.com/gin-gonic/gin"
)

type testSaveRequest struct {
	Title       string `form:"title" binding:"required"`
	Class       string `form:"class" binding:"required"`
	Score       int    `form:"score" binding:"required"`
	TitleType   string `form:"titleType" binding:"required"`
	Difficulty  string `form:"difficulty" binding:"required"`
	Question    string `form:"question" binding:"required"`
	Answer      string `form:"answer" binding:"required"`
}

func GetTestList(c *gin.Context) {
	page := c.Query("page")
	pageInt, err := utils.StringToInt(page)
	if err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	testList, totlePages, err := GetTestLists(pageInt)
	if err != nil {
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  "服务器内部错误: " + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: struct {
			TestLists  []static.TestQuestions
			TitlePages int
		}{TestLists:testList, TitlePages:totlePages},
	})
}

func SaveTest(c *gin.Context)  {
	var req testSaveRequest
	if err := c.Bind(&req); err != nil {
		// 参数解析出错
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误："+err.Error(),
			Data: nil,
		})
		return
	}

	// 获取参数示例
	title := req.Title
	class := req.Class
	score := req.Score
	titleType := req.TitleType
	difficulty := req.Difficulty
	question := req.Question
	answer := req.Answer

	err := SetTest(title,class,score,titleType,difficulty,question,answer)

	if err != nil {
		c.JSON(500,static.Response{
			Code: 10500,
			Msg:  "服务器内部错误："+err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200,static.Response{
		Code: 10200,
		Msg:  "success",
		Data: nil,
	})
}