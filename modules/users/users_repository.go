package users

import (
	"github.com/dee-ex/dx-golang-rest/entities"
)

func (repo *MySQL) GetUserByUsername(username string) (*entities.User, error) {
	var user entities.User
	res := repo.db.Where("username = ?", username).Find(&user)

	return &user, res.Error
}

func (repo *MySQL) CreateNewUser(user *entities.User) (error) {
	res := repo.db.Create(user)

	return res.Error
}

func (repo *MySQL) UpdateToken(username, token string) (error) {
	res := repo.db.Model(&entities.User{}).Where("username = ?", username).Update("token", token)

	return res.Error
}