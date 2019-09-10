package routes

import (
	staticpage "go_simpleweibo/app/controllers/static_page"
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func registerWeb(g *gin.Engine) {
	// ------------------------------ static page ------------------------------
	{

		// 绑定路由 path 和 路由 name，之后可通过 named.G("root") 或 named.GR("root") 获取到路由 path
		// 模板文件中可通过 {{ Route "root" }} 或 {{ RelativeRoute "root" }} 获取 path

		g.GET("/help", staticpage.Help)
		named.Name(g, "help", "GET", "/help")

		g.GET("/about", staticpage.About)
		named.Name(g, "about", "GET", "/about")
	}

}
