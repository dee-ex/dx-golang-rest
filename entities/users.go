package entities

import (
    "time"
)

type (
    User struct {
        ID int
        Username string
        Password string
        DateCreated *time.Time
        Token string
    }
)
