package services

import "server/app/services"

type AccountService struct {

}

func NewAccountService() services.IAccountService {
	return &AccountService{}
}

func (a *AccountService) ValidateUser(username string, password string) (valid bool, error string) {
	return true, ""
}