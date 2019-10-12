package routes

import (

	"go_simpleweibo/middleware"

	ginSessions "github.com/tommy351/gin-sessions"

	"github.com/gin-gonic/gin"
)

var (
	sessionKeyPairs  = []byte("secret123")
	sessionStoreName = "my_session"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {

	g.Use(gin.Recovery())
	g.Use(gin.Logger())

	store := ginSessions.NewCookieStore(sessionKeyPairs)
	store.Options(ginSessions.Options{
		HttpOnly: true,
		Path:     "/",
		MaxAge:   86400 * 30,
	})
	g.Use(ginSessions.Middleware(sessionStoreName, store))

	// 自定义全局中间件
	//g.Use(middleware.Csrf())     // csrf
	//g.Use(middleware.OldValue()) // 记忆上次表单提交的内容，消费即消失
	//g.Use(middleware.GetUser())  // 从 session 中获取用户

	//g.NoRoute(func(c *gin.Context) {
	//	controllers.Render404(c)
	//})
	// web
	registerWeb(g)

	return g
}
