package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const headers string = `Origin, X-Requested-With, Content-Type, Accept, Accept-Language, Authorization, UUID, Token`

/**
设置访问的中间件
*/
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		var url string
			url = "*"
		if origin != "" {
			// 可将将* 替换为指定的域名
			//https://www.huwadayou.com
			c.Header("Access-Control-Allow-Origin", url)
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Header("Access-Control-Allow-Headers", headers)
			c.Header("Access-Control-Max-Age", "86400")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			c.Header("Content-Type", "multipart/form-data;charset=UTF-8")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}