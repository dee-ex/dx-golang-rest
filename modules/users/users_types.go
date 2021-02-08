package users

import (
    "gorm.io/gorm"

    "github.com/dee-ex/dx-golang-rest/entities"
)

type (
    Repository interface {
        GetListUser(offset, limit int) ([]entities.User, error)
        GetUserByID(id int) (*entities.User, error)
        GetUserByUsername(username string) (*entities.User, error)
        CreateNewUser(user *entities.User) (error)
        UpdatePasswordByID(id int, username string) (error)
        UpdateTokenByUsername(username, token string) (error)
        DeleteUserByID(id int) (error)
    }

    MySQL struct {
        db *gorm.DB
    }

    Service struct {
        repo Repository
    }

    UserCreationInput struct {
        Username string
        Password string
    }

    UserLoginInput struct {
        Username string
        Password string
    }

    LoginMessage struct {
        Token string
        TokenLifetime int
    }

    LogoutMessage struct {
        Message string
    }

    ProfileMessage struct {
        ID int
        Username string
    }
)
