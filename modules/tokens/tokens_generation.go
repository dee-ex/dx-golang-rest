package tokens

import (
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

// Create your tokens here

func CreateAuthToken(lifetime int, data map[string]interface{}) (string, error) {
    // Claim the token
    token := jwt.New(jwt.SigningMethodHS256)
    token_claims := token.Claims.(jwt.MapClaims)
    token_claims["authorized"] = true
    token_claims["exp"] = time.Now().Add(time.Duration(lifetime)*time.Minute).Unix()

    // Insert data to claims for token
    for k, v := range data {
        token_claims[k] = v
    }

    // Create the token
    auth_token, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
    if err != nil {
        return "", err
    }

    return auth_token, nil
}
