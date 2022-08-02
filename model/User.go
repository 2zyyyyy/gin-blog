package model

import (
	e "gin-blog/utils/errmsg"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// CheckUser 检查用户是否存在
func CheckUser(name string) int {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return e.ErrorUsernameUsed
	}
	return e.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) int {
	err := db.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetUsers 查询用户列表
func GetUsers(size, num int) []User {
	var users []User
	err := db.Limit(size).Offset((num - 1) * size).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return users
}
