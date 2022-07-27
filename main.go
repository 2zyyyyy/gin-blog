package main

import (
	"gin-blog/model"
	"gin-blog/routers"
)

func main() {
	routers.InitRouter()
	// mysql 初始化
	model.InitDB()
}
