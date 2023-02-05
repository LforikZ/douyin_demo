// @Author Junyi Tan 2023/2/3 23:30:00
package mysql

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string `gorm:"unique"`
	Password      string
	FollowCount   int64
	FollowerCount int64
	IsFollow      bool
}

type UserService struct {
	Name     string `form:"username" json:"username" binding:"required,max=32"`
	Password string `form:"password" json:"password" binding:"required,max=32"`
}

type InfoService struct {
	id   int64  `form:"id" json:"id" binding:"required"`
	Name string `form:"username" json:"username" binding:"required,max=32"`
}

const (
	PassWordCost = 12 //密码加密难度
)

func CreateUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func RegisterAuth(service *UserService, user *User) int64 {
	var count int64
	db.Model(User{}).Where("name=?", service.Name).First(&user).Count(&count)
	return count
}

func LoginAuth(service *UserService, user *User) error {
	err := db.Where("name=?", service.Name).First(&user).Error
	return err
}

func InfoAuth(user *User, id uint) error {
	err := db.Where("id=?", id).First(&user).Error
	return err
}

//SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

//CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
