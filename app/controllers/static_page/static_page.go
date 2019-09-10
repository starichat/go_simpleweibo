package staticpage

import (
	"go_simpleweibo/app/controllers"

	"github.com/gin-gonic/gin"
)

// Help 帮助页
func Help(c *gin.Context) {
	controllers.Render(c, "static_page/help.html", gin.H{})
}

// About 关于页
func About(c *gin.Context) {
	controllers.Render(c, "static_page/about.html", gin.H{})
}
