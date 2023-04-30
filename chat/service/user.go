package service

import (
	"github.com/chat/model"
	"github.com/chat/serializer"
)

type UserRegisterService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=5,max=15"`
	Password string `form:"password" json:"password" binding:"required,min=8,max=16"`
	Email    string `form:"email" json:"email" binding:"required,min=8,max=16"`
}

func (service *UserRegisterService) Register() serializer.Response {
	if len(service.UserName) < 5 {
		return serializer.Response{
			Msg: "用户名至少为5位",
		}
	}
	if len(service.Password) < 8 {
		return serializer.Response{
			Msg: "密码至少为8位",
		}
	}
	if len(service.Email) < 8 {
		return serializer.Response{
			Msg: "邮箱至少为8位",
		}
	}

	var user model.User
	var count1, count2 int64
	model.DB.Model(&model.User{}).Where("email=?", service.Email).Count(&count1)
	if count1 == 1 {
		return serializer.Response{
			Msg: "该邮箱已被注册",
		}
	}
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).Count(&count2)
	if count2 == 1 {
		return serializer.Response{
			Msg: "该用户名已被使用",
		}
	}
	user = model.User{
		UserName: service.UserName,
		Email:    service.Email,
	}
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Response{
			Msg: "加密密码出错",
		}
	}

	if err := model.DB.Create(&user).Error; err != nil {
		return serializer.Response{
			Msg: "创建用户出错",
		}
	}
	return serializer.Response{
		Msg: "注册成功",
	}
}
