package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const USER_COLLECTION_NAME string = "users"

type (
	UserQueries interface {
		CreateUser(ctx context.Context, firstName string, lastName string, email string, password string) (*User, error)
	}

	User struct {
		Id        primitive.ObjectID `bson:"_id"`
		firstName string             `bson:"first_name"`
		lastName  string             `bson:"last_name"`
		email     string             `bson:"email"`
		password  string             `bson:"password"`

		createdAt primitive.DateTime `createdAt:"password"`
		updatedAt primitive.DateTime `updatedAt:"password"`
	}
)

func (u *User) BeforeCreate() error {
	now := primitive.NewDateTimeFromTime(time.Now())

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.password), bcrypt.MaxCost)
	if err != nil {
		return err
	}
	u.password = string(hashedPassword)

	u.createdAt = now
	u.createdAt = now

	return nil
}

func (dbm *mongodb_impl) CreateUser(ctx context.Context, firstName string, lastName string, email string, password string) (*User, error) {

	user := &User{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		password:  password,
	}

	if err := user.BeforeCreate(); err != nil {
		dbm.logger.Error(err.Error())
		return nil, err
	}

	col := dbm.doc.Collection(USER_COLLECTION_NAME)

	_, err := col.InsertOne(ctx, user)

	return user, err
}
