package service

import (
	"context"
	"errors"
	"time"
)

type (
	CreatePostInput struct {
		Title string  `json:"title"`
		Body  string  `json:"body"`
		Image *string `json:"image,omitempty"`
		Video *string `json:"video,omitempty"`
	}

	PostsResponse struct {
		Id        string    `json:"id"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		Slug      string    `json:"slug"`
		Image     *string   `json:"image,omitempty"`
		Video     *string   `json:"video,omitempty"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`

		User *UserResponse `json:"user,omitempty"`
	}

	PostsServices interface {
		CreatePost(input CreatePostInput, userId string) (*PostsResponse, error)
	}
)

func (srv service) CreatePost(input CreatePostInput, userId string) (*PostsResponse, error) {
	post, err := srv.db.AddPost(context.TODO(), input.Title, input.Body, userId, input.Video, input.Image)

	if err != nil {
		srv.logger.Error(err.Error())
		return nil, errors.New("Error while creating posts")
	}

	return &PostsResponse{
		Id:        post.Id.Hex(),
		Title:     post.Title,
		Body:      post.Body,
		Slug:      post.Slug,
		Image:     post.Image,
		Video:     post.Video,
		CreatedAt: post.CreatedAt.Time(),
	}, nil
}
