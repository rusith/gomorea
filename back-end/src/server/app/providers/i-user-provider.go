package providers

import "server/app/models"

type IUserProvider interface {
	GetByUsername(username string) (*models.User, error)
	Add(user *models.User) (*models.User, error)
}
