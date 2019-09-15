package controllers

import (
	"fmt"
	"go_simpleweibo/app/auth"
	"go_simpleweibo/app/helpers"
	viewmodels "go_simpleweibo/app/view_models"
	"go_simpleweibo/config"
	"go_simpleweibo/pkg/flash"
	"go_simpleweibo/routes/named"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	renderObj = map[string]interface{}
)

// Render : 渲染 html
func Render(c *gin.Context, tplPath string, data renderObj) {
	obj := make(renderObj)
    flashStore:=flash.Read(c)
    oldValueStore:=flash.ReadOldFormValue(c)
    validateMsgArr:=flash.ReadValidateMessage(c)
    obj[flash.FlashInContextAndCookieKeyName]=flashStore.Data
	// 上次 post form 的数据，用于回填
	obj[flash.OldValueInContextAndCookieKeyName] = oldValueStore.Data
	// 上次表单的验证信息
	obj[flash.ValidateContextAndCookieKeyName] = validateMsgArr



	// 获取当前登录的用户 (如果用户登录了的话，中间件中会通过 session 存储用户数据)
	if user, err := auth.GetCurrentUserFromContext(c); err == nil {
		obj[config.AppConfig.ContextCurrentUserDataKey] = viewmodels.NewUserViewModelSerializer(user)
	}
	fmt.Println()
	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}
	fmt.Println(data)
	c.HTML(http.StatusOK, tplPath, obj)
}

// RenderError : 渲染错误页面
func RenderError(c *gin.Context, code int, msg string) {
	errorCode := code
	if code == 419 || code == 403 {
		errorCode = 403
	}

	c.HTML(code, "error/error.html", gin.H{
		"errorMsg":  msg,
		"errorCode": errorCode,
		"errorImg":  helpers.Static("/svg/" + strconv.Itoa(code) + ".svg"),
		"backUrl":   named.G("root"),
	})
}

// Render403 -
func Render403(c *gin.Context, msg string) {
	RenderError(c, http.StatusForbidden, msg)
}

// Render404 -
func Render404(c *gin.Context) {
	RenderError(c, http.StatusNotFound, "很抱歉！您浏览的页面不存在。")
}
