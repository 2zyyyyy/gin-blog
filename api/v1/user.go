package v1

import (
	"gin-blog/model"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UserExist 查询用户是否存在
func UserExist(ctx *gin.Context) {

}

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	code := model.CheckUser(data.Username)
	if code == e.SUCCESS {
		// 调用model层数据操作
		model.CreateUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// GetUser 查询单个用户
func GetUser(ctx *gin.Context) {

}

// GetUsers 查询用户列表
func GetUsers(ctx *gin.Context) {
	// 获取前端分页数据
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))
	pageNum, _ := strconv.Atoi(ctx.Query("pageNum"))
	// 默认值
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageSize = -1
	}
	// 调用model层数据操作
	data := model.GetUsers(pageSize, pageNum)
	code := e.SUCCESS
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// EditUser 编辑用户
func EditUser(ctx *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {

}
