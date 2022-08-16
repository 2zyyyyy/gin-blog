package v1

import (
	"gin-blog/middleware"
	"gin-blog/model"
	res "gin-blog/utils"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(ctx *gin.Context) {
	// 获取参数
	user := model.User{}
	if err := ctx.ShouldBindJSON(&user); err != nil {
		res.ResponseErrorWithMsg(ctx, e.ERROR, e.ERROR.GetMsg())
		return
	}
	// 设置token
	token, tokenCode := middleware.SetToken(user.Username)
	// token 生成失败返回错误
	if tokenCode != e.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"code": tokenCode,
			"msg":  "token生成失败",
		})
		return
	}
	// 业务逻辑
	if code := model.LoginCheck(user.Username, user.Password); code == e.SUCCESS {
		// 1.成功
		ctx.JSON(http.StatusOK, gin.H{
			"code":  e.SUCCESS,
			"msg":   e.SUCCESS.GetMsg(),
			"token": token,
		})
		return
	} else if code == e.ErrorUserNotExist {
		// 2.用户不存在
		res.ResponseErrorWithMsg(ctx, code, code.GetMsg())
		return
	} else if code == e.ErrorPasswordWrong {
		// 3.密码错误
		res.ResponseErrorWithMsg(ctx, code, code.GetMsg())
		return
	} else if code == e.ErrorUserNoRight {
		// 4.权限不够
		res.ResponseErrorWithMsg(ctx, code, code.GetMsg())
		return
	}
}
