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
			g1.POST("/login", account.Login)
			g1.POST("/register", account.Register)
		}

		//Browse
		g2 := g.Group("/browse")
		{
			g2.GET("/user",browse.Browse)
			g2.GET("/testList",TestList.GetTestList)
			g2.POST("/saveTest",TestList.SaveTest)
			g2.POST("/deleteTests",TestList.DeleteTest)
			g2.GET("/searchTests",TestList.SearchTest)
		}
	}
	return router
}