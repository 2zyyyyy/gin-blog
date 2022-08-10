package model

import (
	"encoding/base64"
	e "gin-blog/utils/errmsg"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

// UserList 查询用户列表结构体（定义需要返回的字段）
type UserList struct {
	Id        uint
	Username  string
	Role      int
	CreatedAt string
}

// CheckUserByName 检查用户名是否存在（name）
func CheckUserByName(name string) e.ResCode {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return e.ErrorUsernameUsed
	}
	return e.SUCCESS
}

// CheckUserById 检查用户是否存在（id）
func CheckUserById(id int) e.ResCode {
	var count int64
	var user User
	db.Where("id = ?", id).First(&user).Count(&count)
	log.Printf("count:%d\n", count)
	if count == 0 {
		return e.ErrorUserNotExist
	}
	return e.SUCCESS
}

// CheckUpdateUser 更新用户信息 检查用户名是否存在
func CheckUpdateUser(id int, user User) e.ResCode {
	var dbUser User
	// 根据接口入参判断当前用户是否存在
	db.Where("id = ?", id).First(&dbUser)
	// case1:如果用户已删除
	if dbUser.ID == 0 {
		return e.ErrorUserNotExist
	}
	dbUser = User{}
	// case2：非当前用户 无法修改已存在的用户名（对比id和db中查询的id）
	db.Where(&User{Username: user.Username}).Find(&dbUser)
	if user.Username == dbUser.Username && id != int(dbUser.ID) {
		log.Printf("user.username:%s db.username:%s user.id:%d db.id:%d", user.Username, dbUser.Username,
			id, dbUser.ID)
		return e.ErrorUsernameUsed
	}
	// case3：如果查询结果的id和当前修改用户的id相同 则放行
	if id == int(dbUser.ID) {
		return e.SUCCESS
	}
	return e.SUCCESS
}

// CreateUser 新增用户
func CreateUser(data *User) e.ResCode {
	// 对用户密码加密
	//data.Password = ScryptPassword(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// GetUser 查询单个用户
func GetUser(id int) UserList {
	var user UserList
	// 用主键检索
	err := db.Table("user").First(&user, id).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return UserList{}
	}
	return user
}

// GetUsers 查询用户列表
func GetUsers(pageSize, pageNum int) []UserList {
	var userList []UserList
	db.Table("user").Select("id", "username", "role", "created_at").
		Limit(pageSize).Offset((pageNum - 1) * pageSize).Scan(&userList)
	log.Println(userList)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil
	}
	return userList
}

// EditUser 编辑用户
func EditUser(id int, u *User) e.ResCode {
	// 当通过 struct 更新时，GORM 只会更新非零字段。
	// 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	var user = make(map[string]interface{})
	user["username"] = u.Username
	user["role"] = u.Role
	err := db.Model(u).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// DeleteUser 删除用户
func DeleteUser(id int) e.ResCode {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return e.ERROR
	}
	return e.SUCCESS
}

// BeforeSave 使用钩子函数对用户密码加密
func (u *User) BeforeSave(_ *gorm.DB) error {
	u.Password = ScryptPassword(u.Password)
	return nil
}

// ScryptPassword 用户密码加密
func ScryptPassword(password string) string {
	salt := make([]byte, 8)
	salt = []byte{1, 12, 123, 2, 23, 4, 56, 98}

	HashPassword, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, 10)
	if err != nil {
		log.Fatal(err)
	}
	finePassword := base64.StdEncoding.EncodeToString(HashPassword)
	return finePassword
}
