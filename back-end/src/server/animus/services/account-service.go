package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"server/app/providers"
	"server/app/services"
)

type AccountService struct {
	userProvider providers.IUserProvider
}

func NewAccountService(userProvider providers.IUserProvider) services.IAccountService {
	return &AccountService{ userProvider }
}

func (a *AccountService) ValidateUser(username string, password string) (valid bool, error error) {
	user, err := a.userProvider.GetByUsername(username)
	if err != nil { return false, errors.New("invalid username or password")}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return false, errors.New("invalid username or password")
	}

	return true,nil
}