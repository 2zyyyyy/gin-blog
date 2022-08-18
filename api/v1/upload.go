package v1

import (
	"gin-blog/server"
	"gin-blog/utils"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Upload 上传图片接口
func Upload(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("file")
	if err != nil {
		log.Printf("get ctx.Request.FormFile failed, err:%s\n", err)
		utils.ResponseError(ctx, errmsg.ERROR)
	}

	fileSize := fileHeader.Size
	url, code := server.UploadFile(file, fileSize)
	if code == errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  code.GetMsg(),
			"url":  url,
		})
	} else {
		log.Printf("server.UploadFile failed, err:%s\n", err)
		utils.ResponseError(ctx, errmsg.ERROR)
	}
}
