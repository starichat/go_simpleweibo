package user

import (
	"go_simpleweibo/app/helpers"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func sendConfirmEmail(u *userModel.User) error {
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/confirm.html"
	confirmURL := named.G("signup.confirm", "token", u.ActivationToken)

	return helpers.SendMail([]string{u.Email}, subject, tpl, gin.H{"confirmURL": confirmURL})
}
