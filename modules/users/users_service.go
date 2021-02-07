package users

import (
	"time"

	"github.com/dee-ex/dx-golang-rest/entities"
	"github.com/dee-ex/dx-golang-rest/modules/tokens"
)

func (serv *Service) GetUserByUsername(username string) (*entities.User, error) {
	return serv.repo.GetUserByUsername(username)
}

func (serv *Service) CreateNewUser(data UserCreationInput) (*entities.User, error) {
	new_user := entities.User{}

	// Prepare some additional data
	current_time := time.Now()

	new_user.Username = data.Username
	new_user.Password = data.Password
	new_user.DateCreated = &current_time

	err := serv.repo.CreateNewUser(&new_user)
	return &new_user, err
}

func (serv *Service) CreateNewToken(username string) (*LoginMessage, error) {
	login_message := LoginMessage{}

	token_string, err := tokens.CreateAuthToken(username)
	if err != nil {
		return nil, err
	} 

	err = serv.repo.UpdateToken(username, token_string)
	if err != nil {
		return nil, err
	}

	login_message.Token = token_string

	return &login_message, nil
}

func (serv *Service) ClearToken(username string) (*LogoutMessage, error) {
	logout_message := LogoutMessage{}

	err := serv.repo.UpdateToken(username, "")
	if err != nil {
		return nil, err
	}

	logout_message.Message = "Success"

	return &logout_message, nil
}