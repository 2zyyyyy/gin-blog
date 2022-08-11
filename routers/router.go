package routers

import (
	"gin-blog/api/v1"
	"gin-blog/utils"
	"gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.App.AppMode)
	r := gin.Default()

	// 定义路由组
	routerV1 := r.Group("api/v1")
	{
		// 测试路由
		routerV1.GET("test", func(ctx *gin.Context) {
			utils.ResponseSuccess(ctx, errmsg.SUCCESS)
		})
		// 用户相关
		routerV1.POST("user/add", v1.AddUser)      // 添加用户
		routerV1.PUT("user/:id", v1.EditUser)      // 编辑用户
		routerV1.DELETE("user/:id", v1.DeleteUser) // 删除用户

		routerV1.GET("user/:id", v1.GetUser)  // 根据id查询单个用户
		routerV1.GET("user/all", v1.GetUsers) // 查询全部用户

		// 分类相关
		routerV1.POST("category/add", v1.AddCategory)      // 添加分类
		routerV1.PUT("category/:id", v1.EditCategory)      // 编辑分类
		routerV1.DELETE("category/:id", v1.DeleteCategory) // 删除文章

		routerV1.GET("category/:id", v1.GetCategory)   // 根据id查询单个分类
		routerV1.GET("category/all", v1.GetCategories) // 查询全部分类

		// 文章相关
		routerV1.POST("article/add", v1.AddArticle)      // 添加文章
		routerV1.PUT("article/:id", v1.EditArticle)      // 编辑文章
		routerV1.DELETE("article/:id", v1.DeleteArticle) // 删除文章

		routerV1.GET("article/:id", v1.GetArticleDetail)               // 获取文章详情
		routerV1.GET("article/all", v1.GetArticleList)                 // 查询全部文章
		routerV1.GET("article/category/:id", v1.GetArticlesByCategory) // 根据分类查询文章
		//routerV1.GET("article/search", v1.GetArticleList) // 根据标题搜索文章
	}
	// 启动
	_ = r.Run(utils.App.HttpPort)
}
