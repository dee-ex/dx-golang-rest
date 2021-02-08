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

func (serv *Service) CreateNewToken(user *entities.User) (*LoginMessage, error) {
    // Prepare data for token generation
    // Lifetime of a token, on a minute scale
    token_lifetime := 1

    // Your data that you want inject to token
    data := make(map[string]interface{})
    data["id"] = user.ID
    data["username"] = user.Username

    token_string, err := tokens.CreateAuthToken(token_lifetime, data)
    if err != nil {
        return nil, err
    } 

    // After having token string, we update token to user for later access
    err = serv.repo.UpdateTokenByUsername(user.Username, token_string)
    if err != nil {
        return nil, err
    }

    login_message := LoginMessage{}
    login_message.Token = token_string
    login_message.TokenLifetime = token_lifetime

    return &login_message, nil
}

func (serv *Service) ClearToken(username string) (*LogoutMessage, error) {
    err := serv.repo.UpdateTokenByUsername(username, "")
    if err != nil {
        return nil, err
    }

    logout_message := LogoutMessage{}
    logout_message.Message = "Success"

    return &logout_message, nil
}
