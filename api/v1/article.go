package v1

import (
	"gin-blog/model"
	res "gin-blog/utils"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// AddArticle 添加文章
func AddArticle(ctx *gin.Context) {
	var data model.Article
	// 获取数据
	_ = ctx.ShouldBindJSON(&data)
	if code := model.CreateArticle(&data); code == e.SUCCESS {
		res.ResponseSuccess(ctx, nil)
	} else {
		res.ResponseError(ctx, code)
	}
}

// GetArticleDetail 获取文章详情
func GetArticleDetail(ctx *gin.Context) {
	// 获取文章的id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	// 判断当前id的文章是否存在
	if code := model.CheckArticleById(id); code == e.SUCCESS {
		data := model.GetArticleDetail(id)
		res.ResponseSuccess(ctx, data)
	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorArticleNotExist.GetMsg())
	}
}

// GetArticleList 获取文章列表
func GetArticleList(ctx *gin.Context) {
	// 获取分页
	pageSize, pageNum, err := res.PageInfo(ctx)
	if err != nil {
		log.Fatalf("分页数据获取失败, err:%s\n", err)
	}
	// 调用model层数据操作
	data, code := model.GetArticles(pageSize, pageNum)
	if code == e.SUCCESS {
		res.ResponseSuccess(ctx, data)
	} else {
		res.ResponseError(ctx, code)
	}

}

// GetArticlesByCategory 根据分类查询对应文章列表
func GetArticlesByCategory(ctx *gin.Context) {
	// 获取分类id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	// 获取分页
	pageSize, pageNum, err := res.PageInfo(ctx)
	if err != nil {
		log.Fatalf("分页数据获取失败, err:%s\n", err)
	}
	// 调用model层查询数据
	data, code := model.GetArticlesByCategory(id, pageSize, pageNum)
	if code == e.SUCCESS {
		res.ResponseSuccess(ctx, data)
	} else {
		res.ResponseError(ctx, code)
	}
}

// EditArticle 编辑文章
func EditArticle(ctx *gin.Context) {
	// 获取数据
	var data model.Article
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Fatalf("ShouldBindJSON failed, err:%s\n", err)
	}
	// 判断编辑的文章是否存在
	if code := model.CheckArticleById(id); code == e.SUCCESS {
		_ = model.EditArticle(id, &data)
		res.ResponseSuccess(ctx, data)
	} else if code == e.ErrorArticleNotExist {
		// 如果传入的文章id不存在 返回错误
		res.ResponseError(ctx, e.ErrorArticleNotExist)
		return
	} else {
		res.ResponseError(ctx, e.ErrorArticleNotExist)
	}
}

// DeleteArticle 删除文章
func DeleteArticle(ctx *gin.Context) {
	// 获取需要删除的分类id
	id, _ := strconv.Atoi(ctx.Param("id"))
	//判断当前id的分类是否存在
	code := model.CheckArticleById(id)
	if code == e.SUCCESS {
		// 调用model层的数据操作
		code = model.DeleteArticle(id)
		res.ResponseSuccess(ctx, nil)

	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorArticleNotExist.GetMsg())
	}
}
