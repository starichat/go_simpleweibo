package user

import (
	"fmt"
	"go_simpleweibo/app/helpers"
	userModel "go_simpleweibo/app/models/user"
	"go_simpleweibo/routes/named"

	"github.com/gin-gonic/gin"
)

func sendConfirmEmail(u *userModel.User) error {
	fmt.Println("----4.---")
	subject := "感谢注册 Weibo 应用！请确认你的邮箱。"
	tpl := "mail/confirm.html"
	confirmURL := named.G("signup.confirm", "token", u.ActivationToken)
	fmt.Println("-----5.----")

	return helpers.SendMail([]string{u.Email}, subject, tpl, gin.H{"confirmURL": confirmURL})
}
