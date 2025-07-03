package repository

import "github.com/PirateWindows/DriverMockup/internal/model"

var users []model.User

func GetUsers() []model.User {
	return users
}

func AddUser(user model.User) {
	users = append(users, user)
}
