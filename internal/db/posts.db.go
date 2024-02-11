package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

		EditPost(ctx context.Context, id string, title *string, body *string, video *string, image *string, userId string) (*Post, error)
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

func (mdb mongodb_impl) EditPost(ctx context.Context, id string, title *string, body *string, video *string, image *string, userId string) (*Post, error) {
	col := mdb.doc.Collection(POSTS_COLLECTION_NAME)

	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	userIdHex, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key:   "_id",
		Value: idHex,
	}, {
		Key:   "user_id",
		Value: userIdHex,
	}}

	updater := bson.D{}

	if title != nil {
		updater = append(updater, bson.E{Key: "$set", Value: bson.D{{Key: "title", Value: *title}}})
	}

	if body != nil {
		updater = append(updater, bson.E{Key: "$set", Value: bson.D{{Key: "body", Value: *body}}})
	}

	if video != nil {
		updater = append(updater, bson.E{Key: "$set", Value: bson.D{{Key: "video", Value: *video}}})
	}

	if image != nil {
		updater = append(updater, bson.E{Key: "$set", Value: bson.D{{Key: "image", Value: *image}}})
	}

	_, err = col.UpdateOne(ctx, filter, updater)
	if err != nil {
		return nil, err
	}

	post := &Post{}
	res := col.FindOne(ctx, filter)
	err = res.Decode(&post)

	return post, err

}
