package routes

import (
	staticpage "go_simpleweibo/app/controllers/static_page"
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

}
