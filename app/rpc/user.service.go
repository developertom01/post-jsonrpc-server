package rpc

import (
	"context"
	"encoding/json"
	"errors"

	jsonrpc2 "github.com/developertom01/jsonrpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
	"github.com/developertom01/post-jsonrpc-server/utils"
)

type userService struct {
	service service.Service
}

func NewUserService(service service.Service) *userService {
	return &userService{
		service: service,
	}
}

func validateEmailInput(args map[string]any) (*string, []string) {
	errors := make([]string, 0)

	email, ok := args["email"]
	if !ok {
		errors = append(errors, "email does not exist")
	}

	emailString, ok := email.(string)
	if !ok {
		errors = append(errors, "email must be a string")
	}

	ok = utils.ValidateEmail(emailString)
	if !ok {
		errors = append(errors, "Invalid email provided")
	}
	return &emailString, errors
}

func validatePasswordInput(args map[string]any) (*string, []string) {
	errors := make([]string, 0)

	password, ok := args["password"]
	if !ok {
		errors = append(errors, "password does not exist")
	}

	passwordString, ok := password.(string)
	if !ok {
		errors = append(errors, "password must be a string")
	}

	return &passwordString, errors
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

	email, emailErrors := validateEmailInput(args)
	if emailErrors != nil && len(emailErrors) > 0 {
		errors["email"] = emailErrors
	}

	password, passwordError := validatePasswordInput(args)
	if passwordError != nil && len(passwordError) > 0 {
		errors["password"] = passwordError
	}

	if len(errors) > 1 {
		return nil, errors
	}

	return &service.UserCreationInput{
		FirstName: firstNameString,
		Email:     *email,
		LastName:  lastNameString,
		Password:  *password,
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

func ValidateLoginInput(args map[string]any) (*service.LoginInput, map[string][]string) {
	errors := make(map[string][]string)

	email, emailErrors := validateEmailInput(args)

	if emailErrors != nil && len(emailErrors) > 0 {
		errors["email"] = emailErrors
	}

	password, passwordError := validatePasswordInput(args)
	if passwordError != nil && len(passwordError) > 0 {
		errors["password"] = passwordError
	}

	if len(errors) > 1 {
		return nil, errors
	}

	return &service.LoginInput{
		Email:    *email,
		Password: *password,
	}, nil
}

func (srv userService) Login(ctx context.Context, args map[string]any) (*service.LoginResponse, error, *jsonrpc2.RpcErrorCode) {
	input, validationErrors := ValidateLoginInput(args)

	if len(validationErrors) > 1 {
		errStr, _ := json.Marshal(validationErrors)
		code := jsonrpc2.INVALID_PARAMS

		return nil, errors.New(string(errStr)), &code
	}

	loginResponse, err := srv.service.LoginUser(*input)
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR
		return nil, err, &code
	}

	return loginResponse, nil, nil
}

func (srv userService) RefreshToken(ctx context.Context, refreshToken string) (*service.AuthToken, error, *jsonrpc2.RpcErrorCode) {
	tokens, err := srv.service.RefreshToken(refreshToken)
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR
		return nil, err, &code

	}

	return tokens, nil, nil
}
