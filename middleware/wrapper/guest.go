package wrapper

import "github.com/gin-gonic/gin"

func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//TODO: 用户如果已经认证，执行跳转到用户自己的当前页面
		handler(c)
	}
}