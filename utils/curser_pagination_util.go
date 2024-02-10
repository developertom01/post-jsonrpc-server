package utils

import (
	"encoding/base64"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateCursor(id primitive.ObjectID) string {
	return base64.StdEncoding.EncodeToString([]byte(id.Hex()))
}

func GetObjectIdFromCursor(enc string) (primitive.ObjectID, error) {
	objIdString, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return primitive.ObjectIDFromHex(string(objIdString))
}

func GetCursorFilter(cursor *string) (primitive.D, error) {
	filter := primitive.D{}
	if cursor != nil && *cursor != "" {
		cursorObjectId, err := GetObjectIdFromCursor(*cursor)
		if err != nil {
			return nil, err
		}

		filter = bson.D{{Key: "_id", Value: bson.D{{Key: "$gt", Value: cursorObjectId}}}}
	}
	return filter, nil
}

type (
	OrderByDirection int
	OrderBy          struct {
		Field     string
		Direction OrderByDirection
	}
)

const (
	ASC  OrderByDirection = 1
	DESC OrderByDirection = -1
)
