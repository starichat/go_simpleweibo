package password


import (
	"go_simpleweibo/app/controllers"
	passwordResetModel "go_simpleweibo/app/models/password_reset"

	passwordRequest "go_simpleweibo/app/requests/password"
	"go_simpleweibo/pkg/flash"

	"github.com/gin-gonic/gin"

	"fmt"
)


// ShowLinkRequestsForm 显示充值密码的邮箱发送
func ShowLinkRequestsForm(c *gin.Context) {
	fmt.Println("aaaa12132")
	controllers.Render(c, "password/email.html", gin.H{})
}

// SendResetLinkEmail 邮箱发送重设连接
func SendResetLinkEmail(c *gin.Context) {
	fmt.Println("aaaa")
	email := c.PostForm("email")
	passwordForm := &passwordRequest.PasswordEmailForm{
		Email: email,
	}
	pwd, errors := passwordForm.ValidateAndGetToken()

	fmt.Println("-----------:-------", pwd)

	if len(errors) != 0 || pwd == nil {
		flash.SaveValidateMessage(c, errors)
		fmt.Println("aaaabbbb")
		controllers.RedirectRouter(c, "password.request")
		return
	}

	if err := sendResetEmail(pwd); err != nil {
		flash.NewDangerFlash(c, "重置密码邮件发送失败:"+ err.Error())
		passwordResetModel.DeleteByEmail(pwd.Email)
	} else {
		flash.NewSuccessFlash(c, "重置密码已经发送到你邮箱了。请注意查收")
	}

	controllers.RedirectRouter(c, "password.request")
}


// ShowResetForm 密码更新页面
func ShowResetForm(c *gin.Context) {
	token := c.Param("token")
	p, err := passwordResetModel.GetByToken(token)
	if err != nil {
		controllers.Render404(c)
		return
	}

	controllers.Render(c, "password/reset.html", gin.H{
		"token": token,
		"email": p.Email,
	})
}

// Reset 执行密码更新操作
func Reset(c *gin.Context) {
	passwordForm := &passwordRequest.PassWordResetForm{
		Token:                c.PostForm("token"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	user, errors := passwordForm.ValidateAndUpdateUser()

	if len(errors) != 0 || user == nil {
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectRouter(c, "password.reset", "token", c.PostForm("token"))
		return
	}

	flash.NewSuccessFlash(c, "重置密码成功")
	controllers.RedirectRouter(c, "root")
}