package service

import (
	"context"
	"errors"
	"time"

	"github.com/developertom01/post-jsonrpc-server/config"
	"github.com/developertom01/post-jsonrpc-server/internal/logger"
	"github.com/developertom01/post-jsonrpc-server/utils"
	"golang.org/x/crypto/bcrypt"
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

	LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	AuthToken struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}

	LoginResponse struct {
		User  UserResponse `json:"user"`
		Token AuthToken    `json:"token"`
	}
	UserService interface {
		CreateUser(input UserCreationInput) (*UserResponse, error)

		LoginUser(input LoginInput) (*LoginResponse, error)

		RefreshToken(refreshToken string) (*AuthToken, error)
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

func (srv *service) LoginUser(input LoginInput) (*LoginResponse, error) {

	user, err := srv.db.GetUserByEmail(context.TODO(), input.Email)
	if err != nil {
		return nil, errors.New("Email does not exist")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("Wrong password")
	}

	tokens, err := generateTokenPair(user.Id.Hex(), srv.logger)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: *tokens,
		User: UserResponse{
			Id:        user.Id.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.Time(),
		},
	}, nil
}

func (srv service) RefreshToken(refreshToken string) (*AuthToken, error) {
	subj, err := utils.ParseJwtToken(refreshToken, config.REFRESH_TOKEN_SECRET)
	if err != nil {
		srv.logger.Error(err.Error())
		return nil, errors.New("Failed to decode token")
	}

	//Check if user is still in db
	user, err := srv.db.GetUserById(context.TODO(), subj)

	if err != nil {
		srv.logger.Error(err.Error())

		return nil, errors.New("User does not exist")
	}

	return generateTokenPair(user.Id.Hex(), srv.logger)
}

func generateTokenPair(userId string, logger logger.Logger) (*AuthToken, error) {
	refreshToken, err := utils.GenerateJwtToken(userId, config.APP_NAME, config.REFRESH_TOKEN_SECRET, config.REFRESH_TOKEN_DURATION)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("Failed to sign authentication token")
	}

	accessToken, err := utils.GenerateJwtToken(userId, config.APP_NAME, config.ACCESS_TOKEN_SECRET, config.ACCESS_TOKEN_DURATION)
	if err != nil {
		logger.Error(err.Error())
		return nil, errors.New("Failed to sign authentication token")
	}

	return &AuthToken{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}
