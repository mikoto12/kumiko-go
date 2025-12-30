package dao

import (
	"kumiko/internal/model"
	"kumiko/pkg/database"
)

func GetUserList() ([]model.User, error) {
	var users []model.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetUserByID(id string) (*model.User, error) {
	var user model.User
	result := database.DB.First(&user, id)
	return &user, result.Error
}

func CreateUser(user *model.User) error {
	return database.DB.Create(user).Error
}
