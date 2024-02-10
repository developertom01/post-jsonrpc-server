package config

import "os"

var DATABASE_URL = os.Getenv("DATABASE_URL")
var APP_PORT = os.Getenv("APP_PORT")
var DATABASE_NAME = os.Getenv("DATABASE_NAME")

var TEST_DATABASE_NAME = "cute_finds_test_db"
var TEST_DATABASE_URL = "mongodb://root:example@localhost:27018/"
