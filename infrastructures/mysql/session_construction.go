package mysql

import (
    "gorm.io/gorm"

    "github.com/dee-ex/dx-golang-rest/infrastructures/types"
)

func ConstructMySQLSession() (*gorm.DB, error) {
    ses_inf := types.SessionInfor{}

    ses_inf.Username = "root"
    ses_inf.Password = ""
    ses_inf.Host = "127.0.0.1"
    ses_inf.Port = 3306
    ses_inf.ConnectionProtocol = "tcp"
    ses_inf.DatabaseName = "dxrest"
    ses_inf.ParseTime = true
    ses_inf.MaxOpenConns = 10
    ses_inf.ConnMaxLifetime = 15

    return NewMySQLSession(&ses_inf) 
}
