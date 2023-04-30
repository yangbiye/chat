package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName string `gorm:"not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique,not null"`
}

// 加密密码
const (
	PassWordCost        = 12       //密码加密难度
	Active       string = "active" //激活用户
)

// 设置密码
func (user *User) SetPassword(password string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(pwd)
	return nil
}

// 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
