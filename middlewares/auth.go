package middlewares

import (
    "context"
    "net/http"

    "github.com/dee-ex/dx-golang-rest/modules/users"
    "github.com/dee-ex/dx-golang-rest/modules/tokens"
)

func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Get token from header
        token_string := r.Header.Get("token")
        if len(token_string) == 0 {
            http.Error(w, "Unauthorized", 401)
            return
        } 

        // Validate Token
        token, err := tokens.ValidateToken(token_string)
        if err != nil {
            http.Error(w, err.Error(), 401)
            return
        }

        // Extract data from token from token for later validation
        inter_id, err := tokens.ExtractValueFromClaims(token, "id")
        if err != nil {
            http.Error(w, err.Error(), 401)
            return
        }
        inter_username, err := tokens.ExtractValueFromClaims(token, "username")
        if err != nil {
            http.Error(w, err.Error(), 401)
            return
        }

        // Convert interface type
        id := inter_id.(float64)
        username := inter_username.(string)

        // Init repository and service
        repo, err := users.NewMySQLRepository()
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        serv := users.NewService(repo)

        // Last validation stage of token
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

        // If all validate stages passed
        // Insert key-value in context here
        ctx := context.WithValue(r.Context(), "id", id)
        ctx = context.WithValue(ctx, "username", username)

        // Serve next handler with context
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
