package user

import (
	"go_simpleweibo/app/auth"
	"go_simpleweibo/app/controllers"
	"go_simpleweibo/app/models"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/pkg/flash"
	"time"

	"github.com/gin-gonic/gin"
)

func ConfirmEmail(c *gin.Context) {
	token := c.Param("token")

	user, err := userModel.GetByActivationToken(token)
	if user == nil || err != nil {
		controllers.Render404(c)
		return
	}

	// 更新用户
	user.Activated = models.TrueTinyint
	user.ActivationToken = ""
	user.EmailVerifiedAt = time.Now()
	if err = user.Update(false); err != nil {
		flash.NewSuccessFlash(c, "用户激活失败: "+err.Error())
		controllers.RedirectRouter(c, "root")
		return
	}

	auth.Login(c, user)
	flash.NewSuccessFlash(c, "恭喜你，激活成功")
	controllers.RedirectRouter(c, "users.show", user.ID)
}
