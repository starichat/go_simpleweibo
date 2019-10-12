package routes

import (
	"github.com/gin-gonic/gin"
	"go_simpleweibo/middleware/wrapper"
	"go_simpleweibo/routes/named"
)

func registerWeb(g *gin.Engine) {


	// 确认邮箱
	g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
	named.Name(g, "signup.confirm", "GET", "/signup/confirm/:token")
	//  static page(reder page)
	// register and login
	userRouter := g.Group("/users")
	{
		// 注册
		userRouter.POST("",wrapper.Guest(user.Register))
		named.Name(userRouter,"user.register","POST","")
		
	}
}