package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const POSTS_COLLECTION_NAME = "posts"

type (
	Post struct {
		Id        primitive.ObjectID `bson:"_id"`
		Title     string             `bson:"title"`
		Body      string             `bson:"body"`
		UserId    primitive.ObjectID `bson:"user_id"`
		Video     *string            `bson:"video,omitempty"`
		Image     *string            `bson:"image,omitempty"`
		Slug      string             `bson:"slug"`
		CreatedAt primitive.DateTime `bson:"create_at"`
		UpdatedAt primitive.DateTime `bson:"updated_at"`
	}

	PostQuery interface {
		AddPost(ctx context.Context, title string, body string, userId string, video *string, image *string) (*Post, error)
	}
)

func (p *Post) BeforeCreate() {

	p.Id = primitive.NewObjectID()

	now := primitive.NewDateTimeFromTime(time.Now())
	p.CreatedAt = now
	p.UpdatedAt = now

	titelSlug := strings.ReplaceAll(p.Title, " ", "-")
	p.Slug = fmt.Sprintf("%s-%s", titelSlug, p.Id.Hex())
}

func (dbm *mongodb_impl) AddPost(ctx context.Context, title string, body string, userId string, video *string, image *string) (*Post, error) {
	userObjectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	post := Post{
		Title:  title,
		Body:   body,
		UserId: userObjectId,
		Video:  video,
		Image:  image,
	}
	post.BeforeCreate()

	col := dbm.doc.Collection(POSTS_COLLECTION_NAME)
	_, err = col.InsertOne(ctx, &post)

	return &post, err
}
