package mysql

import (
    "fmt"
    "time"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"

    "github.com/dee-ex/dx-golang-rest/infrastructures/types"
)

func NewMySQLSession(ses_inf *types.SessionInfor) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=%t",
                        ses_inf.Username,
                        ses_inf.Password,
                        ses_inf.ConnectionProtocol,
                        ses_inf.Host,
                        ses_inf.Port,
                        ses_inf.DatabaseName,
                        ses_inf.ParseTime)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
      return nil, err
    }

    mysql_db, err := db.DB()
    if err != nil {
      return nil, err
    }

    mysql_db.SetMaxOpenConns(ses_inf.MaxOpenConns)
    mysql_db.SetConnMaxLifetime(time.Duration(ses_inf.ConnMaxLifetime)*time.Second)

    return db, nil
}
