package rpc

import (
	"context"
	"encoding/json"
	"errors"

	jsonrpc2 "github.com/developertom01/jsonrpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
	"github.com/developertom01/post-jsonrpc-server/internal/db"
	"github.com/developertom01/post-jsonrpc-server/utils"
)

type postService struct {
	service service.Service
}

func NewPostService(service service.Service) *postService {
	return &postService{
		service: service,
	}
}

func validateCreatePostInput(input map[string]any) (*service.CreatePostInput, map[string][]string) {
	errors := make(map[string][]string)

	title := utils.ValidateRequiredStringField("title", input, errors)

	body := utils.ValidateRequiredStringField("body", input, errors)

	image := utils.ValidateUrlField("image", input, errors)

	video := utils.ValidateUrlField("video", input, errors)

	if len(errors) > 0 {
		return nil, errors
	}

	return &service.CreatePostInput{
		Title: *title,
		Body:  *body,
		Image: image,
		Video: video,
	}, errors
}

func (ps *postService) CreatePost(ctx context.Context, input map[string]any) (*service.PostsResponse, error, *jsonrpc2.RpcErrorCode) {
	user, err := ps.service.HandleAuthorization(ctx)
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR

		return nil, err, &code
	}

	postInput, validationErrors := validateCreatePostInput(input)
	if len(validationErrors) > 0 {
		errorsJson, _ := json.Marshal(validationErrors)
		code := jsonrpc2.INVALID_PARAMS

		return nil, errors.New(string(errorsJson)), &code
	}

	post, err := ps.service.CreatePost(*postInput, user.Id.Hex())
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR

		return nil, errors.New("Failed to Create Post"), &code
	}

	return post, nil, nil
}

func validateStatusField(input map[string]any, errors map[string][]string) *db.PostStatus {
	status, ok := input["status"]
	if !ok {
		return nil
	}

	statusString, ok := status.(string)
	if !ok {
		errors["status"] = append(errors["status"], "Invalid post status")
		return nil
	}

	postStatus := db.PostStatus(statusString)

	return &postStatus
}

func validateEditPostInput(input map[string]any) (*service.EditPostInput, map[string][]string) {
	errors := make(map[string][]string)

	title := utils.ValidateStringField("title", input, errors)

	body := utils.ValidateStringField("body", input, errors)

	image := utils.ValidateUrlField("image", input, errors)

	video := utils.ValidateUrlField("video", input, errors)

	status := validateStatusField(input, errors)

	return &service.EditPostInput{
		Title:  title,
		Body:   body,
		Image:  image,
		Video:  video,
		Status: status,
	}, errors
}

func (ps *postService) EditPost(ctx context.Context, id string, input map[string]any) (*service.PostsResponse, error, *jsonrpc2.RpcErrorCode) {
	user, err := ps.service.HandleAuthorization(ctx)
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR

		return nil, err, &code
	}

	editPostInput, valErrors := validateEditPostInput(input)
	if len(valErrors) > 0 {
		errorsJson, _ := json.Marshal(valErrors)
		code := jsonrpc2.INVALID_PARAMS

		return nil, errors.New(string(errorsJson)), &code
	}

	post, err := ps.service.EditPost(id, *editPostInput, user.Id.Hex())
	if err != nil {
		code := jsonrpc2.INTERNAL_ERROR
		return nil, err, &code
	}

	return post, nil, nil
}
