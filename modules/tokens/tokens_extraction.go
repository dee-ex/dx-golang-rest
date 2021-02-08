package tokens

import (
    "fmt"
    
    "github.com/dgrijalva/jwt-go"
)

func ExtractValueFromClaims(token *jwt.Token, key string) (interface{}, error) {
    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return "", fmt.Errorf("Cannot extract `%s` from token", key)
    }

    value, ok := claims[key]
    if !ok {
        return "", fmt.Errorf("Cannot extract `%s` from token", key)
    }

    return value, nil
}
