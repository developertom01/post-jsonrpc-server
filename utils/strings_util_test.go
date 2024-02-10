package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	var (
		str         = "TestingString"
		expectedStr = "testing_string"
	)
	snakeCase := ToSnakeCase(str)
	assert.Equal(t, expectedStr, snakeCase)
}
