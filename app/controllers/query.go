package controllers

import (
	"github.com/gin-gonic/gin"
	"math"
	"strconv"
)

func GetPageQuery(c *gin.Context, defaultPageLine, totalCount int) (offset, limit, currentPage, pageTotalCount int) {
	page, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		page = 1
	}
	currentPage = page
	pageline, err := strconv.Atoi(c.Query("pageline"))
	if err != nil {
		pageline = defaultPageLine
	}
	page = page - 1

	if page == 0 {
		offset = 0
	} else {
		offset = page * pageline
	}
	limit = pageline

	pageTotalCount = int(math.Ceil(float64(totalCount) / float64(pageline)))
	if pageTotalCount <= 0 {
		pageTotalCount = 1
	}

	return
}
