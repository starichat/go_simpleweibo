package routes

import (
	"go_simpleweibo/app/controllers"

	"github.com/gin-gonic/gin"
)

var (
	sessionKeyPairs  = []byte("secret123")
	sessionStoreName = "my_session"
)

// Register 注册路由和中间件
func Register(g *gin.Engine) *gin.Engine {

	g.NoRoute(func(c *gin.Context) {
		controllers.Render404(c)
	})
	// web
	registerWeb(g)

	return g
}
