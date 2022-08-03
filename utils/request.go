package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PageInfo(ctx *gin.Context) (pageSize, pageNum int, err error) {
	// 处理分页
	size := ctx.Query("pageSize")
	num := ctx.Query("pageNum")
	pageSize, err = strconv.Atoi(size)
	if err != nil {
		pageSize = 10
	}
	pageNum, err = strconv.Atoi(num)
	if err != nil {
		pageNum = 1
	}
	return
}
