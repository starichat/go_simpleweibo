package routes

import (
	"go_simpleweibo/app/controllers/blog"
	"go_simpleweibo/app/controllers/sessions"
	"go_simpleweibo/app/controllers/followers"
	staticpage "go_simpleweibo/app/controllers/static_page"
	"go_simpleweibo/app/controllers/password"
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
			// 展示具体用户页面
			userRouter.GET("/show/:id", wrapper.Auth(user.Show))
			named.Name(userRouter, "users.show", "GET", "/show/:id")
			// 关注用户
			userRouter.POST("/followers/store/:id", wrapper.Auth(followers.Store))
			named.Name(userRouter, "followers.store", "POST", "/followers/store/:id")
			
		}

	}

	// session
	{
		// 登录界面
		g.GET("/login", wrapper.Guest(sessions.Create))
		named.Name(g, "login.create", "GET", "/login")
		// 登录
		g.POST("/login", wrapper.Guest(sessions.Store))
		named.Name(g, "login.store", "POST", "/login")

		// 注销登陆
		g.POST("/logout", sessions.Destroy)
		named.Name(g, "login.destroy", "POST", "/logout")
		named.Name(g, "logout", "POST", "/logout")
	}

	// 重置密码
	passwordRouter := g.Group("/password")
	{
		//显示重置密码的邮箱发送页面
		passwordRouter.GET("/reset",wrapper.Guest(password.ShowLinkRequestsForm))
		named.Name(passwordRouter, "password.request", "GET", "/reset")
		// 邮箱发送重置链接
		passwordRouter.POST("/email", wrapper.Guest(password.SendResetLinkEmail))
		named.Name(passwordRouter, "password.email", "POST", "/email")
		// 密码更新页面
		passwordRouter.GET("/reset/:token", wrapper.Guest(password.ShowResetForm))
		named.Name(passwordRouter, "password.reset", "GET", "/reset/:token")
		// 执行密码更新操作
		passwordRouter.POST("/reset", wrapper.Guest(password.Reset))
		named.Name(passwordRouter, "password.update", "POST", "/reset")
	}

	// blog
	blogRouter := g.Group("/blogs")
	{
		// 发布博客
		blogRouter.POST("",wrapper.Auth(blog.Store))
		named.Name(blogRouter,"blogs.store","POST","")
		// 处理删除微博请求
		blogRouter.POST("/destroy/:id",wrapper.Auth(blog.Destroy))
		named.Name(blogRouter,"blogs.destroy","POST","/destroy")

	}

}
