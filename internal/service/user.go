package service

import (
	"kumiko/internal/dao"
	"kumiko/internal/model"
)

func GetUserList() ([]model.User, error) {
	return dao.GetUserList()
}

func GetUserByID(id string) (interface{}, error) {
	return dao.GetUserByID(id)
}

func CreateUser(user *model.User) error {
	return dao.CreateUser(user)
}
