package routers

import (
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
		}
	}
	return router
}