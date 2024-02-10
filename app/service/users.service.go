package service

import (
	"context"
	"time"
)

type (
	UserCreationInput struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	UserResponse struct {
		Id        string    `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
	}

	UserService interface {
		CreateUser(input UserCreationInput) (*UserResponse, error)
	}
)

func (srv *service) CreateUser(input UserCreationInput) (*UserResponse, error) {
	user, err := srv.db.CreateUser(context.TODO(), input.FirstName, input.LastName, input.Email, input.Password)

	if err != nil {
		return nil, err
	}

	return &UserResponse{
		Id:        user.Id.Hex(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time(),
	}, nil
}
