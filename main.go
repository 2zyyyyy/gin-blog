package main

import (
	"gin-blog/model"
	"gin-blog/routers"
)

func main() {
	// mysql 初始化
	model.InitDB()
	routers.InitRouter()
}
