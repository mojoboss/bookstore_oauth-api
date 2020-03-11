package rest

import (
	"bytes"
	"encoding/json"
	"github.com/mojoboss/bookstore_users-api/domain/users"
	"github.com/mojoboss/bookstore_users-api/utils/errors"
	"net/http"
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct{}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.LoginRequest{
		Email:    email,
		Password: password,
	}
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, errors.NewInternalServerError("Error in Login user marshalling request")
	}
	resp, err := http.Post("http://localhost:8080", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.NewInternalServerError("Error in Login user while making post request")
	}
	defer resp.Body.Close()
	var user users.User
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return nil, errors.NewInternalServerError("Error in Login user response decode")
	}
	return &user, nil
}
