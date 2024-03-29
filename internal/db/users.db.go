package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

const USER_COLLECTION_NAME string = "users"

type (
	UserQueries interface {
		CreateUser(ctx context.Context, firstName string, lastName string, email string, password string) (*User, error)

		GetUserByEmail(ctx context.Context, email string) (*User, error)
		GetUserById(ctx context.Context, id string) (*User, error)
	}

	User struct {
		Id        primitive.ObjectID `bson:"_id"`
		FirstName string             `bson:"first_name"`
		LastName  string             `bson:"last_name"`
		Email     string             `bson:"email"`
		Password  string             `bson:"password"`

		CreatedAt primitive.DateTime `bson:"created_at"`
		UpdatedAt primitive.DateTime `bson:"updated_at"`
	}
)

func (u *User) BeforeCreate() error {
	u.Id = primitive.NewObjectID()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	now := primitive.NewDateTimeFromTime(time.Now())
	u.CreatedAt = now
	u.UpdatedAt = now

	return nil
}

func (dbm *mongodb_impl) CreateUser(ctx context.Context, firstName string, lastName string, email string, password string) (*User, error) {
	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
	}

	if err := user.BeforeCreate(); err != nil {
		dbm.logger.Error(err.Error())
		return nil, err
	}

	col := dbm.doc.Collection(USER_COLLECTION_NAME)

	_, err := col.InsertOne(ctx, user)

	return user, err
}

func (dbm mongodb_impl) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	col := dbm.doc.Collection(USER_COLLECTION_NAME)
	filter := bson.D{{
		Key:   "email",
		Value: email,
	}}

	user := &User{}

	res := col.FindOne(ctx, filter)
	err := res.Decode(&user)

	return user, err
}

func (dbm mongodb_impl) GetUserById(ctx context.Context, id string) (*User, error) {
	col := dbm.doc.Collection(USER_COLLECTION_NAME)
	idValue, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.D{{
		Key:   "_id",
		Value: idValue,
	}}

	user := &User{}

	res := col.FindOne(ctx, filter)
	err = res.Decode(&user)

	return user, err

}
