package services

import (
	"Go-API-Gin/source/entities"
)

func CreateUser(user entities.User) (*entities.User, error) {
	return &user, nil
}
