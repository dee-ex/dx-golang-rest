package users

import (
    "github.com/dee-ex/dx-golang-rest/infrastructures/mysql"
)

func NewService(repo Repository) *Service {
    return &Service{repo: repo}
}

func NewMySQLRepository() (*MySQL, error) {
    db, err := mysql.ConstructMySQLSession()
    if err != nil {
        return nil, err
    }

    return &MySQL{db: db}, nil
}
