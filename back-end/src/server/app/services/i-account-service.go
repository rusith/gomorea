package services


type IAccountService interface {
	ValidateUser(username string, password string) (valid bool, error error)
}