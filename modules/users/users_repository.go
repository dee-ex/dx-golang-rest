package users

import (
    "github.com/dee-ex/dx-golang-rest/entities"
)

func (repo *MySQL) GetListUser(offset, limit int) ([]entities.User, error) {
    var users []entities.User
    res := repo.db.Offset(offset).Limit(limit).Find(&users)

    return users, res.Error
}

func (repo *MySQL) GetUserByID(id int) (*entities.User, error) {
    var user entities.User
    res := repo.db.Where("id = ?", id).Find(&user)

    return &user, res.Error
}

func (repo *MySQL) GetUserByUsername(username string) (*entities.User, error) {
    var user entities.User
    res := repo.db.Where("username = ?", username).Find(&user)

    return &user, res.Error
}

func (repo *MySQL) CreateNewUser(user *entities.User) (error) {
    res := repo.db.Create(user)

    return res.Error
}

func (repo *MySQL) UpdatePasswordByID(id int, password string) (error) {
    res := repo.db.Model(&entities.User{}).Where("id = ?", id).Update("password", password)

    return res.Error
}

func (repo *MySQL) UpdateTokenByUsername(username, token string) (error) {
    res := repo.db.Model(&entities.User{}).Where("username = ?", username).Update("token", token)

    return res.Error
}

func (repo *MySQL) DeleteUserByID(id int) (error) {
    res := repo.db.Where("id = ?", id).Delete(&entities.User{})

    return res.Error
}
