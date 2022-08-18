package server

import (
	"context"
	"gin-blog/utils"
	"gin-blog/utils/errmsg"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"log"
	"mime/multipart"
)

func UploadFile(file multipart.File, size int64) (string, errmsg.ResCode) {
	putPolicy := storage.PutPolicy{
		Scope: utils.QiNiu.Bucket,
	}
	log.Printf("key:%s secret:%s\n", utils.QiNiu.AccessKey, utils.QiNiu.SecretKey)
	mac := qbox.NewMac(utils.QiNiu.AccessKey, utils.QiNiu.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: false,
		UseHTTPS:      false,
	}

	putExtra := storage.PutExtra{}

	formUploader := storage.NewFormUploader(&cfg)
	res := storage.PutRet{}

	err := formUploader.PutWithoutKey(context.Background(), &res, upToken, file, size, &putExtra)
	if err != nil {
		log.Printf("formUploader.PutWithoutKey failed, err%s\n", err)
		return "", errmsg.ERROR
	}
	url := utils.QiNiu.QiuNiuServer + res.Key
	return url, errmsg.SUCCESS
}
