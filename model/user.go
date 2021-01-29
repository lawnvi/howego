package model

import (
	"github.com/jinzhu/gorm"
	"howego/config/database"
	"howego/config/log"
	"strconv"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Field string `json:"field"`
}

type UserMapper interface {
	Find(email string) User
	Append(user *User)
	Update(user *User)
}

type userDao struct {
	db *gorm.DB
}

func NewUserMapper(db *gorm.DB) UserMapper {
	database.HasTable(User{})
	return &userDao{db: db}
}

func (m *userDao) Append(user *User) {
	m.db.Create(&user)
	log.I("db", "append user ok"+user.Name)
}

func (m *userDao) Find(email string) User {
	var user = User{}
	m.db.Where("email= ?", email).Find(&user)
	return user
}

func (m *userDao) Update(user *User) {
	log.I("db", "update user ok "+strconv.Itoa(user.Id))
}
