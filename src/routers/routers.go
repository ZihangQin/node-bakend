package routers

import (
	"bk/src/handler/TestList"
	"bk/src/handler/account"
	"bk/src/handler/browse"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine{
	router := gin.Default()

	router.Use(Cors())

	g := router.Group("/api")
	{
		//Login
		g1 := g.Group("/account")
		{
			//登录
			g1.POST("/login", account.Login)
			//注册
			g1.POST("/register", account.Register)
		}

		//Browse
		g2 := g.Group("/browse")
		{
			//主页获取用户信息
			g2.GET("/user",browse.Browse)
			//试题管理页面获取试题列表
			g2.GET("/testList",TestList.GetTestList)
			//新增试题
			g2.POST("/saveTest",TestList.SaveTest)
			//删除试题
			g2.POST("/deleteTests",TestList.DeleteTest)
			//搜索试题
			g2.GET("/searchTests",TestList.SearchTest)
		}
	}
	return router
}