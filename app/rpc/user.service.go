package rpc

import (
	"context"
	"encoding/json"
	"errors"

	jsonrpc2 "github.com/developertom01/json-rpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
)

type userService struct {
	service service.Service
}

func NewUserService(service service.Service) *userService {
	return &userService{
		service: service,
	}
}

func ValidateUserInput(args map[string]any) (*service.UserCreationInput, map[string][]string) {
	errors := make(map[string][]string)
	firstName, ok := args["firstName"]
	if !ok {
		errors["firstName"] = []string{"firstName does not exist"}
	}

	firstNameString, ok := firstName.(string)
	if !ok {
		errors["firstName"] = append(errors["firstName"], "firstName must be a string")
	}

	lastName, ok := args["lastName"]
	if !ok {
		errors["lastName"] = []string{"lastName does not exist"}
	}

	lastNameString, ok := lastName.(string)
	if !ok {
		errors["lastName"] = append(errors["lastName"], "lastName must be a string")
	}

	email, ok := args["email"]
	if !ok {
		errors["email"] = []string{"email does not exist"}
	}

	emailString, ok := email.(string)
	if !ok {
		errors["email"] = append(errors["email"], "email must be a string")
	}

	password, ok := args["password"]
	if !ok {
		errors["password"] = []string{"password does not exist"}
	}

	passwordString, ok := password.(string)
	if !ok {
		errors["password"] = append(errors["password"], "password must be a string")
	}

	return &service.UserCreationInput{
		FirstName: firstNameString,
		Email:     emailString,
		LastName:  lastNameString,
		Password:  passwordString,
	}, errors
}

func (srv userService) CreateUser(ctx context.Context, args map[string]any) (*service.UserResponse, error, *jsonrpc2.RpcErrorCode) {

	userCreationInput, validationErrors := ValidateUserInput(args)
	if len(validationErrors) != 0 {
		errorString, _ := json.Marshal(validationErrors)
		var code = jsonrpc2.INVALID_PARAMS
		return nil, errors.New(string(errorString)), &code
	}
	user, err := srv.service.CreateUser(*userCreationInput)
	if err != nil {
		var code = jsonrpc2.INTERNAL_ERROR
		return nil, err, &code
	}

	return user, nil, nil
}
