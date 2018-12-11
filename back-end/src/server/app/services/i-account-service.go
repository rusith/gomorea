package services

import "server/app/models"

type IAccountService interface {
	ValidateUser(username string, password string) (user *models.User, error error)
	CreateUser(user *models.User) (*models.User, error)
}