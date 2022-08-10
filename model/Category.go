package model

import (
	e "gin-blog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// CheckCategoryByName 检查分类名称是否存在（name）
func CheckCategoryByName(name string) e.ResCode {
	var category Category
	db.Select("id").Where("name = ?", name).First(&category)
	if category.ID > 0 {
		return e.ErrorCategoryNameUsed
	}
	return e.SUCCESS
}

// CheckCategoryById 检查分类是否存在（id）
func CheckCategoryById(id int) e.ResCode {
	var count int64
	var category Category
	db.Where("id = ?", id).First(&category).Count(&count)
	log.Printf("count:%d\n", count)
	if count == 0 {
		return e.ErrorCategoryNotExist
	}
	return e.SUCCESS
}

// CheckUpdateCategory 更新分类信息 检查分类名称是否存在
func CheckUpdateCategory(id int, category Category) e.ResCode {
	var dbCategory Category
	// 根据接口入参判断当前分类是否存在
	db.Where("id = ?", id).First(&dbCategory)
	// case1:如果分类不存在
	if dbCategory.ID == 0 {
		return e.ErrorCategoryNotExist
	}
	// 重置dbCategory结构体的值（否则在查询的时候会出现该结构体字段的所有条件）
	dbCategory = Category{}
	// case2：非当前分类 无法修改已存在的分类名（对比id和db中查询的id）
	db.Where(&Category{Name: category.Name}).Find(&dbCategory)
	log.Printf("data.id:%d, db.id:%d, data.name:%s, db.name:%s\n", id, dbCategory.ID, category.Name, dbCategory.Name)
	if category.Name == dbCategory.Name && id != int(dbCategory.ID) {
		return e.ErrorCategoryNameUsed
	}
	// case3：如果查询结果的id和当前修改分类的id相同 则放行
	if id == int(dbCategory.ID) {
		return e.SUCCESS
	}
	return e.SUCCESS
}

// CreateCategory 新增分类
func CreateCategory(data *Category) e.ResCode {
	err := db.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetCategory 查询单个分类
func GetCategory(id int) Category {
	var category Category
	// 用主键检索
	err := db.Table("category").First(&category, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Category{}
	}
	return category
}

// GetCategories 查询分类列表
func GetCategories(pageSize, pageNum int) []Category {
	var categories []Category
	db.Table("category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&categories)
	log.Println(categories)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return categories
}

// EditCategory 编辑分类
func EditCategory(id int, category *Category) e.ResCode {
	var cate = make(map[string]interface{})
	log.Printf("id:%d\n", id)
	cate["name"] = category.Name
	err := db.Model(category).Where("id = ?", id).Updates(category).Select("name").Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteCategory 删除分类
func DeleteCategory(id int) e.ResCode {
	var category Category
	err := db.Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
