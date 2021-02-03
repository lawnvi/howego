package service

import (
	"fmt"
	"howego/config"
	"howego/config/database"
	"howego/model"
)

type UserService struct {
	userDao model.UserDao
}

func NewUserService() UserService {
	return UserService{userDao: model.NewUserDao(database.DB)}
}

func (u *UserService) SignUp(user model.User) config.ResultMsg{
	msg := config.ResultMsg{}
	um := model.NewUserDao(database.DB)
	if um.FindByEmail(user.Email).Id > 0{
		fmt.Printf("user %v is already exist", user)
		msg.Data = "user is already exist"
		msg.Code = 10
	}else {
		um.Append(&user)
		msg.Data = user
	}
	return msg
}

func (u *UserService) GetUserInfo(email string) config.ResultMsg{
	um := model.NewUserDao(database.DB)
	msg := config.ResultMsg{}
	var user = um.FindByEmail(email)
	if user == (model.User{}) {
		msg.Code = 0001
		msg.Msg = "not find this email:"+ email
	}else {
		msg.Data = user
	}
	return msg
}