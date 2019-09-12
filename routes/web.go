package routes

import (
	"go_simpleweibo/app/controllers/sessions"
	staticpage "go_simpleweibo/app/controllers/static_page"
	"go_simpleweibo/app/controllers/user"
	"go_simpleweibo/middleware/wrapper"
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{

		g.GET("/", staticpage.Home)
		named.Name(g, "root", "GET", "/")

		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "GET", "/help")

		g.GET("/about", staticpage.About)
		named.Name(g, "about", "GET", "/about")
	}

	// user
	{
		g.GET("/signup", wrapper.Guest(user.Create))
		named.Name(g, "signup", "GET", "/signup")

		// 注册
		g.GET("/signup/confirm/:token", wrapper.Guest(user.ConfirmEmail))
		// 带参路由绑定，可通过 named.G("signup.confirm", token) 或 named.GR("signup.confirm", token) 获取 path
		// 模板文件中可通过 {{ Route "signup.confirm" .token }} 或 {{ RelativeRoute "signup.confirm" .token }} 获取 path
		named.Name(g, "signup.confirm", "GET", "/signup/confirm/:token")

		userRouter := g.Group("/users")
		{
			// 创建用户页面
			userRouter.GET("/create", wrapper.Guest(user.Create))
			named.Name(userRouter, "user.create", "GET", "/create")
			// 保存新用户
			userRouter.POST("", wrapper.Guest((user.Store)))
			named.Name(userRouter, "users.store", "POST", "")
		}

	}

	// session
	{
		// 登录界面
		g.GET("/login", wrapper.Guest(sessions.Create))
		named.Name(g, "login.create", "GET", "/login")
		// 登录
		g.POST("/login", wrapper.Guest(sessions.Store))
		named.Name(g, "login", "POST", "/login")
	}

}
