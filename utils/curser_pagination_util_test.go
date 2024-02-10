package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateCursor(t *testing.T) {
	var (
		objectId       = "5fcb2f17ae1d3d5d94407037"
		expectedCursor = "NWZjYjJmMTdhZTFkM2Q1ZDk0NDA3MDM3"
	)

	objectIdFromHex, err := primitive.ObjectIDFromHex(objectId)
	assert.Nil(t, err)

	cursor := CreateCursor(objectIdFromHex)
	assert.Equal(t, expectedCursor, cursor)

}

func TestGetObjectIdFromCursor(t *testing.T) {
	var (
		expectedObjectId = "5fcb2f17ae1d3d5d94407037"
		cursor           = "NWZjYjJmMTdhZTFkM2Q1ZDk0NDA3MDM3"
	)

	objectId, err := GetObjectIdFromCursor(cursor)
	assert.Nil(t, err)

	assert.Equal(t, expectedObjectId, objectId.Hex())

}

func TestGetCursorFilter(t *testing.T) {
	id, err := primitive.ObjectIDFromHex("5fcb2f17ae1d3d5d94407037")
	assert.Nil(t, err)

	var cursor = "NWZjYjJmMTdhZTFkM2Q1ZDk0NDA3MDM3"

	expectedBson := bson.D{{Key: "_id", Value: bson.D{{Key: "$gt", Value: id}}}}

	actualBson, err := GetCursorFilter(&cursor)

	assert.Nil(t, err)
	assert.Equal(t, expectedBson, actualBson)
}
