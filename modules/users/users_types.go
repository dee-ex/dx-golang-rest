package users

import (
	"gorm.io/gorm"

	"github.com/dee-ex/dx-golang-rest/entities"
)

type (
	Repository interface {
		GetUserByUsername(username string) (*entities.User, error)
		CreateNewUser(user *entities.User) (error)
		UpdateToken(username, token string) (error)
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
	}

	LogoutMessage struct {
		Message string
	}
)