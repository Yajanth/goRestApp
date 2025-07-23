package service

import (
	"github.com/Yajanth/goRestApp/model"
)

var users = []model.User{
	{ID: 1, Name: "John Doe", Email: "john@example.com"},
}

func GetAllUsers() []model.User {
	return users
}

func AddUser(user model.User) {
	user.ID = len(users) + 1
	users = append(users, user)

}
