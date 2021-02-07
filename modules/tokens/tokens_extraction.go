package tokens

import (
    "fmt"
    
	"github.com/dgrijalva/jwt-go"
)

func ExtractUsernameFromClaims(token *jwt.Token) (string, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", fmt.Errorf("Cannot extract `username` from token")
    }

    username, ok := claims["username"].(string)
    if !ok {
        return "", fmt.Errorf("Cannot extract username from token")
    }

    return username, nil
}