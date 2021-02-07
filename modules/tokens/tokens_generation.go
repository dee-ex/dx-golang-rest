package tokens

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateAuthToken(username string) (string, error) {
	// Claim the token
    token := jwt.New(jwt.SigningMethodHS256)
    token_claims := token.Claims.(jwt.MapClaims)
    token_claims["authorized"] = true
    token_claims["username"] = username
    token_claims["exp"] = time.Now().Add(1*time.Minute).Unix()

    // Create the token
    auth_token, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }

    return auth_token, nil
}