package v1

import (
	"gin-blog/model"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

var code int

// UserExist 查询用户是否存在
func UserExist(ctx *gin.Context) {

}

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	_ = ctx.ShouldBindJSON(&data)
	code = model.CheckUser(data.Username)
	if code == e.ErrorUsernameUsed {
		code = e.ErrorUsernameUsed
	}
	if code == e.SUCCESS {
		model.CreateUser(&data)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  e.GetMsg(code),
	})
}

// GetUser 查询单个用户
func GetUser(ctx *gin.Context) {

}

// GetUsers 查询用户列表
func GetUsers(ctx *gin.Context) {

}

// EditUser 编辑用户
func EditUser(ctx *gin.Context) {

}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {

}
