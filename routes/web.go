package routes

import (
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{
		g.GET("/", staticpage.Home)
		// 绑定路由 path 和 路由 name，之后可通过 named.G("root") 或 named.GR("root") 获取到路由 path
		// 模板文件中可通过 {{ Route "root" }} 或 {{ RelativeRoute "root" }} 获取 path
		named.Name(g, "root", "GET", "/")

		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "GET", "/help")

		g.GET("/about", staticpage.About)
		named.Name(g, "about", "GET", "/about")
	}

}
