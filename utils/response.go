package utils

import (
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 封装返回结果

type ResponseData struct {
	Code e.ResCode   `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"` // 忽略null
}

func ResponseError(ctx *gin.Context, code e.ResCode) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  code.GetMsg(),
		Data: nil,
	})
}

func ResponseSuccess(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code: e.SUCCESS,
		Msg:  e.SUCCESS.GetMsg(),
		Data: data,
	})
}

func ResponseErrorWithMsg(ctx *gin.Context, code e.ResCode, msg interface{}) {
	ctx.JSON(http.StatusOK, &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	})
}
