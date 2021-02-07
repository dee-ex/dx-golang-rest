package users

import (
	"net/http"
	"encoding/json"

	"github.com/dee-ex/dx-golang-rest/utils/responses"
)

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	// Get input data
	var data UserCreationInput
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Some basic validations
	if len(data.Username) == 0 {
		http.Error(w, "No `username` found!", 400)
		return
	}
	if len(data.Password) == 0 {
		http.Error(w, "No `password` found!", 400)
		return
	}

	// `username` uniqueness validation stage

	// Init repository and service
	repo, err := NewMySQLRepository()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	serv := NewService(repo)

	temp_user, err := serv.GetUserByUsername(data.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if temp_user.ID != 0 {
		http.Error(w, "`username` is already exists.", 409)
		return	
	}

	// Validation stage passed
	new_user, err := serv.CreateNewUser(data)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Fetch the result
	responses.JSONResponse(w, 200, new_user)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get input data
	var data UserCreationInput
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Some basic validations
	if len(data.Username) == 0 {
		http.Error(w, "No `username` found!", 400)
		return
	}
	if len(data.Password) == 0 {
		http.Error(w, "No `password` found!", 400)
		return
	}

	// Init repository and service
	repo, err := NewMySQLRepository()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	serv := NewService(repo)

	user, err := serv.GetUserByUsername(data.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return	
	}
	if user.ID == 0 {
		http.Error(w, "Uncorrect username.", 404)
		return
	}

	if user.Password != data.Password {
		http.Error(w, "Password does not match", 401)
		return
	}

	// Login stage passed
	// Create token for later access
	login_message, err := serv.CreateNewToken(data.Username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Fetch the result
	responses.JSONResponse(w, 200, login_message)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Init repository and service
	repo, err := NewMySQLRepository()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	serv := NewService(repo)

	// Clear the token
	logout_message, err := serv.ClearToken(r.Context().Value("username").(string))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Fetch the result
	responses.JSONResponse(w, 200, logout_message)
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	responses.TextResponse(w, 200, r.Context().Value("username").(string))
}