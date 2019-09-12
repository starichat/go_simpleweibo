package user

import (
	"fmt"
	"go_simpleweibo/app/controllers"
	userRequest "go_simpleweibo/app/requests/user"
	"go_simpleweibo/pkg/flash"

	"github.com/gin-gonic/gin"
)

// Create 创建用户页面
func Create(c *gin.Context) {
	fmt.Println("--------------------------create user------------------------------")
	controllers.Render(c, "user/create.html", gin.H{})
}

// post 提交注册用户信息
func Store(c *gin.Context) {
	// 验证参数和创建用户
	userCreateForm := &userRequest.UserCreateForm{
		Name:                 c.PostForm("name"),
		Email:                c.PostForm("email"),
		Password:             c.PostForm("password"),
		PasswordConfirmation: c.PostForm("password_confirmation"),
	}
	fmt.Println("--------------------------user:---------------------------", userCreateForm.Name)
	fmt.Println("--------------------------user:---------------------------", userCreateForm.Email)
	fmt.Println("--------------------------user:---------------------------", userCreateForm.Password)
	fmt.Println("--------------------------user:---------------------------", userCreateForm.PasswordConfirmation)

	user, errors := userCreateForm.ValidateAndSave()

	if len(errors) != 0 || user == nil {
		fmt.Println("----------------------------ERROR------------------------------------")
		flash.SaveValidateMessage(c, errors)
		controllers.RedirectRouter(c, "users.create")
		return
	}

	if err := sendConfirmEmail(user); err != nil {
		fmt.Println("--------------------------email ERROR------------------------------------")
		flash.NewDangerFlash(c, "验证邮件发送失败: "+err.Error())
	} else {
		flash.NewSuccessFlash(c, "验证邮件已发送到你的注册邮箱上，请注意查收。")
	}

	controllers.RedirectRouter(c, "root")
}
