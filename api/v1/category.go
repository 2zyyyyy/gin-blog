package v1

import (
	"gin-blog/model"
	res "gin-blog/utils"
	e "gin-blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// AddCategory 添加分类
func AddCategory(ctx *gin.Context) {
	var data model.Category
	// 获取数据
	_ = ctx.ShouldBindJSON(&data)
	if code := model.CheckCategoryByName(data.Name); code == e.SUCCESS {
		_ = model.CreateCategory(&data)
		res.ResponseSuccess(ctx, nil)
	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorCategoryNameUsed.GetMsg())
	}
}

// GetCategory 查询单个分类信息
func GetCategory(ctx *gin.Context) {
	// 获取查询分类的id
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	// 判断当前id的分类是否存在
	if code := model.CheckCategoryById(id); code == e.SUCCESS {
		data := model.GetCategory(id)
		res.ResponseSuccess(ctx, data)
	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorCategoryNotExist.GetMsg())
	}
}

// GetCategories 查询分类列表
func GetCategories(ctx *gin.Context) {
	// 获取分页
	pageSize, pageNum, err := res.PageInfo(ctx)
	if err != nil {
		log.Fatalf("分页数据获取失败, err:%s\n", err)
	}
	data := model.GetCategories(pageSize, pageNum)
	res.ResponseSuccess(ctx, data)
}

// EditCategory 编辑分类
func EditCategory(ctx *gin.Context) {
	// 获取数据
	var data model.Category
	id, err := strconv.Atoi(ctx.Param("id"))
	log.Println(id)
	if err != nil {
		log.Fatalf("ctx.Param failed, err:%s\n", err)
	}
	err = ctx.ShouldBindJSON(&data)
	if err != nil {
		log.Fatalf("ShouldBindJSON failed, err:%s\n", err)
	}
	// 判断编辑的分类是否存在
	if code := model.CheckUpdateCategory(id, data); code == e.SUCCESS {
		_ = model.EditCategory(id, &data)
		res.ResponseSuccess(ctx, nil)
	} else if code == e.ErrorCategoryNotExist {
		// 如果传入的分类id不存在 返回错误
		res.ResponseError(ctx, e.ErrorCategoryNotExist)
		return
	} else {
		res.ResponseError(ctx, e.ErrorCategoryNameUsed)
	}
}

// DeleteCategory 删除分类
func DeleteCategory(ctx *gin.Context) {
	// 获取需要删除的分类id
	id, _ := strconv.Atoi(ctx.Param("id"))
	//判断当前id的分类是否存在
	if code := model.CheckCategoryById(id); code == e.SUCCESS {
		// 调用model层的数据操作
		code = model.DeleteCategory(id)
		res.ResponseSuccess(ctx, nil)
	} else {
		res.ResponseErrorWithMsg(ctx, code, e.ErrorCategoryNotExist.GetMsg())
	}
}
