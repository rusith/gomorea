package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/app/models"
	"server/app/providers"
	"server/app/services"
)

type AccountService struct {
	userProvider providers.IUserProvider
}

func NewAccountService(userProvider providers.IUserProvider) services.IAccountService {
	return &AccountService{ userProvider }
}

func (a *AccountService) ValidateUser(username string, password string) (valid *models.User, error error) {
	user, err := a.userProvider.GetByUsername(username)
	if err != nil { return nil, errors.New("invalid username or password")}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, errors.New("invalid username or password")
	}

	return user,nil
}

func (a *AccountService) CreateUser( user *models.User) (*models.User, error) {
	pwdBytes := []byte(user.Password)
	pwd, enErr := bcrypt.GenerateFromPassword(pwdBytes, bcrypt.MinCost)
	if enErr != nil { return nil, enErr }
	user.Password = string(pwd[:])
	u, e := a.userProvider.Add(user)
	return u, e
}