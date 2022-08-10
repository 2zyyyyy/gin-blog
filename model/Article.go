package model

import (
	e "gin-blog/utils/errmsg"
	"gorm.io/gorm"
	"log"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model
	Title        string `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int    `gorm:"type:int;not null" json:"cid"`
	Desc         string `gorm:"type:varchar(200)" json:"desc"`
	Content      string `gorm:"type:longtext" json:"content"`
	Img          string `gorm:"type:varchar(100)" json:"img"`
	CommentCount int    `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int    `gorm:"type:int;not null;default:0" json:"read_count"`
}

// CheckArticleById 检查文章是否存在（id）
func CheckArticleById(id int) e.ResCode {
	var count int64
	var article Article
	db.Where("id = ?", id).First(&article).Count(&count)
	if count == 0 {
		return e.ErrorArticleNotExist
	}
	return e.SUCCESS
}

// CreateArticle 新增分文章
func CreateArticle(data *Article) e.ResCode {
	err := db.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetArticleDetail 文章详情
func GetArticleDetail(id int) Article {
	var article Article
	// 用主键检索
	err := db.Table("article").First(&article, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Article{}
	}
	return article
}

// GetArticles 查询文章列表
func GetArticles(pageSize, pageNum int) []Article {
	var articles []Article
	db.Table("article").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articles)
	log.Println(articles)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return articles
}

// EditArticle 编辑分类
func EditArticle(id int, article *Article) e.ResCode {
	var maps = make(map[string]interface{})
	log.Printf("id:%d\n", id)
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["content"] = article.Content
	maps["img"] = article.Img

	err := db.Model(article).Where("id = ?", id).Updates(&maps).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteArticle 删除文章
func DeleteArticle(id int) e.ResCode {
	var article Article
	err := db.Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}
