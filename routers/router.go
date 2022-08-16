package routers

import (
	"gin-blog/api/v1"
	"gin-blog/middleware"
	"gin-blog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.App.AppMode)
	r := gin.Default()

	auth := r.Group("api/v1")
	// 分组设置鉴权中间件
	auth.Use(middleware.JwtMiddleware())
	// 需要token鉴权的接口
	{
		// user
		auth.PUT("user/:id", v1.EditUser)      // 编辑用户
		auth.DELETE("user/:id", v1.DeleteUser) // 删除用户

		// category
		auth.POST("category/add", v1.AddCategory)      // 添加分类
		auth.PUT("category/:id", v1.EditCategory)      // 编辑分类
		auth.DELETE("category/:id", v1.DeleteCategory) // 删除文章

		// article
		auth.POST("article/add", v1.AddArticle)      // 添加文章
		auth.PUT("article/:id", v1.EditArticle)      // 编辑文章
		auth.DELETE("article/:id", v1.DeleteArticle) // 删除文章

	}
	// 无需token鉴权
	routerV1 := r.Group("api/v1")
	{
		// login
		routerV1.POST("login", v1.Login) // 登录接口
		// user
		routerV1.POST("user/add", v1.AddUser) // 用户注册
		routerV1.GET("user/:id", v1.GetUser)  // 根据id查询单个用户
		routerV1.GET("user/all", v1.GetUsers) // 查询全部用户

		// category
		routerV1.GET("category/:id", v1.GetCategory)   // 根据id查询单个分类
		routerV1.GET("category/all", v1.GetCategories) // 查询全部分类

		// article
		routerV1.GET("article/:id", v1.GetArticleDetail)               // 获取文章详情
		routerV1.GET("article/all", v1.GetArticleList)                 // 查询全部文章
		routerV1.GET("article/category/:id", v1.GetArticlesByCategory) // 根据分类查询文章
		//routerV1.GET("article/search", v1.GetArticleList) // 根据标题搜索文章
	}
	// 启动
	_ = r.Run(utils.App.HttpPort)
}
