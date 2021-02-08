package types

type (
    SessionInfor struct {
        Username string
        Password string
        Host string
        Port int
        ConnectionProtocol string
        DatabaseName string
        ParseTime bool
        MaxOpenConns int
        ConnMaxLifetime int
    }
)
