package routers

import (
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	gin.SetMode(utils.App.AppMode)
	r := gin.Default()

	// 定义路由zu
	router := r.Group("api/v1")
	{
		router.GET("hello", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 1000,
				"msg":  "success",
			})
		})
	}
	// 启动
	_ = r.Run(utils.App.HttpPort)
}
