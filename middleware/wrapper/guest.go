// 只有非登录用户才能访问
package wrapper

import "github.com/gin-gonic/gin"

func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 用户已经登录则跳转到 root page
		currentUser,err:auth
	}
}
