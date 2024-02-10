package config

import (
	"os"
	"time"
)

var DATABASE_URL = os.Getenv("DATABASE_URL")
var APP_PORT = os.Getenv("APP_PORT")
var APP_NAME = os.Getenv("APP_NAME")
var DATABASE_NAME = os.Getenv("DATABASE_NAME")

var TEST_DATABASE_NAME = "cute_finds_test_db"
var TEST_DATABASE_URL = "mongodb://root:example@localhost:27018/"

var REFRESH_TOKEN_SECRET = os.Getenv("REFRESH_TOKEN_SECRET")
var ACCESS_TOKEN_SECRET = os.Getenv("ACCESS_TOKEN_SECRET")
var REFRESH_TOKEN_DURATION = time.Hour * 24
var ACCESS_TOKEN_DURATION = time.Minute * 15
