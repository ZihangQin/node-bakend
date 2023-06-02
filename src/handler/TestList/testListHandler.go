package TestList

import (
	"bk/src/static"
	"bk/src/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type testSaveRequest struct {
	Title      string `form:"content" binding:"required" json:"content"`
	Class      string `form:"grade" binding:"required" json:"grade"`
	Score      string `form:"score" binding:"required" json:"score"`
	TitleType  string `form:"type" binding:"required" json:"type"`
	Difficulty string `form:"difficulty" binding:"required" json:"difficulty"`
	Answer     string `form:"answer" binding:"required" json:"answer"`
	Token      string `json:"token"`
}


type testUpdateRequest struct {
	Id string `json:"id"`
	Types string `json:"type"`
	Grade string `json:"grade"`
	Content string `json:"content"`
	Difficulty string `json:"difficulty"`
	Score string `json:"score"`
	Answer string `json:"answer"`
	Token string `json:"token"`
}

func GetTestList(c *gin.Context) {
	page := c.Query("page")
	token := c.Query("token")
	pageInt, err := utils.StringToInt(page)
	if err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误",
			Data: nil,
		})
		return
	}
	testList, totlePages, err := GetTestLists(pageInt, token)
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
			TestLists  interface{}
			TitlePages int
		}{TestLists: testList, TitlePages: totlePages},
	})
}

func SaveTest(c *gin.Context) {
	var req testSaveRequest
	if err := c.Bind(&req); err != nil {
		// 参数解析出错
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误：" + err.Error(),
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
	answer := req.Answer
	token := req.Token

	err := SetTest(title, class, score, titleType, difficulty, answer, token)

	if err != nil {
		fmt.Println(err)
		if err.Error() == "参数不可为空" {
			c.JSON(200, static.Response{
				Code: 10201,
				Msg:  err.Error(),
				Data: nil,
			})
		}
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  "服务器内部错误：" + err.Error(),
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: nil,
	})
}

func DeleteTest(c *gin.Context) {
	var data struct {
		StrList map[string]string `json:"strList"`
	}
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数格式错误",
			Data: nil,
		})
		return
	}
	err := DeleteTests(data.StrList)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  "数据库删除错误",
			Data: nil,
		})
		return
	}

	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: nil,
	})
}

func SearchTest(c *gin.Context) {
	data := c.Query("data")
	fmt.Println(data)
	tests, totle, _ := SearchTests(data)
	if len(tests) <= 0 {
		c.JSON(200, static.Response{
			Code: 10201,
			Msg:  "success",
			Data: nil,
		})
		return
	}
	c.JSON(200, static.Response{
		Code: 10200,
		Msg:  "success",
		Data: struct {
			Test  interface{} `json:"test"`
			Totle int         `json:"totle"`
		}{Test: tests, Totle: totle},
	})
}

func UpdateTest(c *gin.Context)  {
	var req testUpdateRequest
	if err := c.Bind(&req); err != nil {
		// 参数解析出错
		c.JSON(400, static.Response{
			Code: 10400,
			Msg:  "参数错误：" + err.Error(),
			Data: nil,
		})
		return
	}

	err := UpdateTests(req.Id,req.Types,req.Grade,req.Content,req.Difficulty,req.Score,req.Answer,req.Token)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, static.Response{
			Code: 10500,
			Msg:  err.Error(),
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