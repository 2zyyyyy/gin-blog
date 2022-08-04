package v1

import (
	"gin-blog/model"
	res "gin-blog/utils"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// AddUser 添加用户
func AddUser(ctx *gin.Context) {
	var data model.User
	// 获取数据
	_ = ctx.ShouldBindJSON(&data)
	code := model.CheckUserByName(data.Username)
	if code == e.SUCCESS {
		// 调用model层数据操作
		model.CreateUser(&data)
		res.ResponseSuccess(ctx, data)
	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorUsernameUsed.GetMsg())
	}

}

// GetUser 查询单个用户
func GetUser(ctx *gin.Context) {

}

// GetUsers 查询用户列表
func GetUsers(ctx *gin.Context) {
	// 获取分页
	pageSize, pageNum, err := res.PageInfo(ctx)
	if err != nil {
		log.Fatalf("分页数据获取失败, err:%s\n", err)
	}
	// 调用model层数据操作
	data := model.GetUsers(pageSize, pageNum)
	res.ResponseSuccess(ctx, data)
}

// EditUser 编辑用户
func EditUser(ctx *gin.Context) {
	// 获取数据
	var data model.User
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Fatalf("ShouldBindJSON failed, err:%s\n", err)
	}
	// 判断编辑的用户是否存在
	code := model.CheckUpdateUser(id, data)
	if code == e.ErrorUserNotExist {
		// 如果传入的用户id不存在 返回错误
		res.ResponseError(ctx, e.ErrorUserNotExist)
		return
	}
	// 判断修改后的用户名是否存在
	code = model.CheckUserByName(data.Username)
	if code == e.SUCCESS {
		// 如果不存在 操作model层
		model.EditUser(int(data.ID), &data)
		res.ResponseSuccess(ctx, data)
	} else {
		// 如果存在 返回错误
		res.ResponseError(ctx, e.ErrorUsernameUsed)
	}
}

// DeleteUser 删除用户
func DeleteUser(ctx *gin.Context) {
	// 获取需要删除的用户id
	id, _ := strconv.Atoi(ctx.Param("id"))
	//判断当前id的用户是否存在
	code := model.CheckUserById(id)
	if code == e.SUCCESS {
		// 调用model层的数据操作
		code = model.DeleteUser(id)
		res.ResponseSuccess(ctx, nil)

	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorUserNotExist.GetMsg())
	}
}
