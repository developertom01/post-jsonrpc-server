package rpc

import (
	"context"
	"encoding/json"
	"errors"

	jsonrpc2 "github.com/developertom01/json-rpc2"
	"github.com/developertom01/post-jsonrpc-server/app/service"
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

	title, ok := input["title"]
	if !ok {
		errors["title"] = []string{"title is required"}
	}

	titleString, ok := title.(string)
	if !ok {
		errors["title"] = append(errors["title"], "title must be a string")
	}

	body, ok := input["body"]
	if !ok {
		errors["body"] = []string{"body is required"}
	}
	bodyString, ok := body.(string)
	if !ok {
		errors["body"] = append(errors["body"], "body must be a string")
	}

	image, ok := input["image"]
	var imageString *string
	if ok {
		imageStr, ok := image.(string)
		if !ok {
			errors["image"] = []string{"image must be a string"}
		}

		ok = utils.ValidateUrl(imageStr)

		if !ok {
			errors["image"] = append(errors["image"], "image must be a valid url")
		}
		imageString = &imageStr
	}

	video, ok := input["video"]
	var videoString *string
	if ok {
		videoStr, ok := video.(string)
		if !ok {
			errors["video"] = []string{"video must be a string"}
		}

		ok = utils.ValidateUrl(videoStr)

		if !ok {
			errors["video"] = append(errors["video"], "video must be a valid url")
		}
		videoString = &videoStr
	}

	return &service.CreatePostInput{
		Title: titleString,
		Body:  bodyString,
		Image: imageString,
		Video: videoString,
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
