package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-users-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

// Repository interface specifies the needed methods for a users repository
type Repository interface {
	Create(user *User) (*User, *utils.RestErr)
	Retrive(id string) (*User, *utils.RestErr)
	Delete(id string) *utils.RestErr
	Update(id, email, picture, fullName, givenName, familyName, description string) *utils.RestErr
	AddGroup(id, groupID string) *utils.RestErr
	DelGroup(id, groupID string) *utils.RestErr
	AddContact(id, contactID string) *utils.RestErr
	DelContact(id, contactID string) *utils.RestErr
}

type repository struct{}

// NewRepository returns a new repository that implements the Repository interface.
func NewRepository() Repository {
	r := repository{}
	return &r
}

func (r *repository) Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := usersC.InsertOne(ctx, user)
	if err != nil {
		return nil, utils.BadRequest("can't create user.")
	}
	return user, nil
}

func (r *repository) Retrive(id string) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := usersC.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, utils.BadRequest("can't operate or find the user.")
	}
	return &user, nil
}

func (r *repository) Delete(id string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result, err := usersC.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return utils.InternalServerErr("can't operate delete functionality.")
	}
	if result.DeletedCount == 0 {
		return utils.NotFound("user doesn't exist.")
	}
	return nil
}

func (r *repository) Update(id, email, picture, fullName, givenName, familyName, description string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$set": bson.M{
			"email":       email,
			"picture":     picture,
			"fullName":    fullName,
			"givenName":   givenName,
			"familyName":  familyName,
			"description": description,
		},
	}
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate update functionality.")
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("nothing to update user is already up-to-date.")
	}
	return nil
}

func (r *repository) AddGroup(id, groupID string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$addToSet": bson.M{
			"JoinedGroupIDs": groupID,
		},
	}
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate add group functionality.")
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("user is already a member of the group.")
	}
	return nil
}

func (r *repository) DelGroup(id, groupID string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$pull": bson.M{
			"JoinedGroupIDs": groupID,
		},
	}
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate del group functionality.")
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("user is not a member of this group.")
	}
	return nil
}

func (r *repository) AddContact(id, contactID string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$addToSet": bson.M{
			"ContactIDs": contactID,
		},
	}
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate add contact functionality.")
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("this contact already exists.")
	}
	return nil
}

func (r *repository) DelContact(id, contactID string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	update := bson.M{
		"$pull": bson.M{
			"ContactIDs": contactID,
		},
	}
	result, err := usersC.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return utils.InternalServerErr("can't operate del contact functionality.")
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("contact is not in contacts list.")
	}
	return nil
}
