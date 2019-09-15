package blog

import (
	"github.com/gin-gonic/gin"
	"go_simpleweibo/app/controllers"
	userModel "go_simpleweibo/app/models/user"
)

// blog 的相关工具类

func backTo(c *gin.Context, currentUser *userModel.User) {
	back := c.DefaultPostForm("back", "")
	if back != "" {
		controllers.Redirect(c, back, true)
		return
	}

	controllers.RedirectRouter(c, "user.show", currentUser.ID)
}
