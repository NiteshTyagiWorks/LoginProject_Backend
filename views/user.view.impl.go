package views

import (
	"LoginProject_Backend/models"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserViewImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewUserView(usercollection *mongo.Collection, ctx context.Context) UserView {
	return &UserViewImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func (u *UserViewImpl) CheckUser(email *string) error {
	var user *models.User
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return err
}

func (u *UserViewImpl) MakeUser(email *string, name *string, picture *string) error {
	var user models.User
	user.Email = *email
	user.Name = *name
	user.Picture = *picture
	user.Username = (*email)[:len(*email)-10]
	user.Password = user.Username
	user.Job = "Worker"
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
}

func (u *UserViewImpl) GetUser(email *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "email", Value: email}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
}

func (u *UserViewImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "email", Value: user.Email}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{
		bson.E{Key: "username", Value: user.Username},
		bson.E{Key: "name", Value: user.Name},
		bson.E{Key: "password", Value: user.Password},
		bson.E{Key: "job", Value: user.Job},
	}}}
	result, _ := u.usercollection.UpdateOne(u.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no such user found")
	}
	return nil
}

func (u *UserViewImpl) LoginUser(username *string, password *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "username", Value: username}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		query = bson.D{bson.E{Key: "email", Value: username}}
		err = u.usercollection.FindOne(u.ctx, query).Decode(&user)

		if user == nil {
			return user, errors.New("no such user found")
		}
		if user.Password != *password {
			return user, errors.New("wrong password")
		}
	}
	if user.Password != *password {
		return user, errors.New("wrong password")
	}

	return user, err
}
