// 只有非登录用户才能访问
package wrapper

import (
	"fmt"
	"go_simpleweibo/app/auth"
	"go_simpleweibo/app/controllers"
	"go_simpleweibo/pkg/flash"

	"github.com/gin-gonic/gin"
)	

func Guest(handler gin.HandlerFunc) gin.HandlerFunc {
	fmt.Println("login 1")
	return func(c *gin.Context) {
		// 用户已经登录则跳转到 root page
		currentUser, err := auth.GetCurrentUserFromContext(c)
		if currentUser != nil || err == nil {
			fmt.Println("----------currentuser:",currentUser)
			flash.NewInfoFlash(c, "您已登录，无需再次操作！")
			controllers.RedirectRouter(c, "root")
			fmt.Println("handle(c)")
			return
		}

		fmt.Println("handle(c)")

		handler(c)
	}
}
