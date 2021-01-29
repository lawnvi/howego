package service

import (
	"fmt"
	"howego/config"
	"howego/config/database"
	"howego/model"
)

type UserService interface {
	SignUp (user model.User) config.ResultMsg
	GetUserInfo(email string) config.ResultMsg
}

type userImpl struct {
}

func NewUserService() UserService {
	return &userImpl{}
}

func (u *userImpl) SignUp(user model.User) config.ResultMsg{
	msg := config.ResultMsg{}
	um := model.NewUserMapper(database.DB)
	if um.Find(user.Email).Id > 0{
		fmt.Printf("user %v is already exist", user)
		msg.Data = "user is already exist"
		msg.Code = 10
	}else {
		um.Append(&user)
		msg.Data = user
	}
	return msg
}

func (u *userImpl) GetUserInfo(email string) config.ResultMsg{
	um := model.NewUserMapper(database.DB)
	msg := config.ResultMsg{}
	msg.Data = um.Find(email)
	return msg
}