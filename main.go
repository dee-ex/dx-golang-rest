package main

import (
    "os"
    "log"
    "net/http"

    "github.com/rs/cors"
    "github.com/gorilla/mux"
    "github.com/gorilla/handlers"

    "github.com/dee-ex/dx-golang-rest/middlewares"
    "github.com/dee-ex/dx-golang-rest/modules/users"
)

func ServeHTTP() {
    r := mux.NewRouter()

    r.HandleFunc("/register", users.RegistrationHandler).Methods("POST")
    r.HandleFunc("/login", users.LoginHandler).Methods("POST")

    auth_r := r.PathPrefix("").Subrouter()
    auth_r.Use(middlewares.AuthMiddleware)

    auth_r.HandleFunc("/me", users.TestHandler)
    auth_r.HandleFunc("/logout", users.LogoutHandler).Methods("GET")

    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{
                                http.MethodGet,
                                http.MethodPost,
                                http.MethodPut,
                                http.MethodDelete,
                                http.MethodOptions,
                            },
    })

    enhanced_r := handlers.LoggingHandler(os.Stdout, c.Handler(r))

    log.Fatal(http.ListenAndServe(":6969", enhanced_r))
}

func main() {
    // Some configs here
    os.Setenv("SECRET_KEY", "99999999969")
    
    ServeHTTP()
}