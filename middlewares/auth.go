package middlewares

import (
	"context"
	"net/http"

	"github.com/dee-ex/dx-golang-rest/modules/users"
	"github.com/dee-ex/dx-golang-rest/modules/tokens"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token_string := r.Header.Get("token")
        if len(token_string) == 0 {
        	http.Error(w, "Unauthorized", 401)
        	return
        } 

        token, err := tokens.ValidateToken(token_string)
        if err != nil {
        	http.Error(w, err.Error(), 401)
            return
        }

        username, err := tokens.ExtractUsernameFromClaims(token)
        if err != nil {
        	http.Error(w, err.Error(), 401)
            return
        }

        // Init repository and service
        repo, err := users.NewMySQLRepository()
        if err != nil {
        	http.Error(w, err.Error(), 500)
        	return
        }
        serv := users.NewService(repo)

        user, err := serv.GetUserByUsername(username)
        if err != nil {
        	http.Error(w, err.Error(), 500)
            return
        }
        if user.ID == 0 {
        	http.Error(w, "Token may be fake or old.", 401)
        	return
        }

        if user.Token != token_string {
        	http.Error(w, "Token maybe old.", 401)
        	return
        }

        ctx := context.WithValue(r.Context(), "username", username)
        //ctx = context.WithValue(ctx, "another_key", another_value)

        next.ServeHTTP(w, r.WithContext(ctx))
    })
}