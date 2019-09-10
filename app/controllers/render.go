package controllers

import (
	"go_simpleweibo/app/helpers"
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

	// 填充传递进来的数据
	for k, v := range data {
		obj[k] = v
	}

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
