package service

import (
	"context"
	"errors"
	"time"

	"github.com/developertom01/post-jsonrpc-server/internal/db"
)

type (
	CreatePostInput struct {
		Title string  `json:"title"`
		Body  string  `json:"body"`
		Image *string `json:"image,omitempty"`
		Video *string `json:"video,omitempty"`
	}

	EditPostInput struct {
		Title  *string        `json:"title,omitempty"`
		Body   *string        `json:"body,omitempty"`
		Image  *string        `json:"image,omitempty"`
		Video  *string        `json:"video,omitempty"`
		Status *db.PostStatus `json:"status,omitempty"`
	}

	PostsResponse struct {
		Id        string        `json:"id"`
		Title     string        `json:"title"`
		Body      string        `json:"body"`
		Slug      string        `json:"slug"`
		Image     *string       `json:"image,omitempty"`
		Video     *string       `json:"video,omitempty"`
		Status    db.PostStatus `json:"status"`
		CreatedAt time.Time     `json:"created_at"`
		UpdatedAt time.Time     `json:"updated_at"`

		User *UserResponse `json:"user,omitempty"`
	}

	PostsServices interface {
		CreatePost(input CreatePostInput, userId string) (*PostsResponse, error)

		EditPost(id string, input EditPostInput, userId string) (*PostsResponse, error)
	}
)

func (srv *service) CreatePost(input CreatePostInput, userId string) (*PostsResponse, error) {
	post, err := srv.db.AddPost(context.TODO(), input.Title, input.Body, userId, input.Video, input.Image)
	if err != nil {
		srv.logger.Error(err.Error())
		return nil, errors.New("Error while creating post")
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

func (srv *service) EditPost(id string, input EditPostInput, userId string) (*PostsResponse, error) {
	post, err := srv.db.EditPost(context.TODO(), id, input.Title, input.Body, input.Video, input.Image, input.Status, userId)
	if err != nil {
		srv.logger.Error(err.Error())
		return nil, errors.New("Error while updating post")
	}

	return &PostsResponse{
		Id:        post.Id.Hex(),
		Title:     post.Title,
		Body:      post.Body,
		Slug:      post.Slug,
		Status:    post.Status,
		Image:     post.Image,
		Video:     post.Video,
		CreatedAt: post.CreatedAt.Time(),
	}, nil

}
