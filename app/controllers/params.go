package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
)
// 从path参数中获取到int参数
func GetIntParam(c *gin.Context, key string) (int, error) {
	i,err:=strconv.Atoi(c.Param(key))
	if err != nil{
		return 0,err
	}

	return i,nil
}