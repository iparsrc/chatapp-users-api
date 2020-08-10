package domain

import (
	"context"
	"time"

	"github.com/parsaakbari1209/ChatApp-users-api/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Create(user *User) (*User, *utils.RestErr) {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	_, err := usersC.InsertOne(ctx, user)
	if err != nil {
		return nil, utils.InternalServerErr("Can't create user.")
	}
	return user, nil
}

func Retrive(id string, private bool) (*User, *utils.RestErr) {
	var user User
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := usersC.FindOne(ctx, bson.M{"_id": id}).Decode(&user); err != nil {
		return nil, utils.BadRequest(err.Error())
	}
	// At this point the 'user' is filled with all the user information on database.
	if private {
		// Return user's private profile.
		return &user, nil
	}
	// Remove user's private data.
	user.Email = ""
	user.FullName = ""
	user.FamilyName = ""
	user.DateCreated = ""
	user.ContactsIDs = []string{}
	user.JoinedGroupIDs = []string{}
	// Return user's public profile.
	return &user, nil
}

func Delete(id string) *utils.RestErr {
	usersC := db.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	result, err := usersC.DeleteMany(ctx, bson.M{"_id": id})
	if err != nil {
		return utils.BadRequest(err.Error())
	}
	if result.DeletedCount == 0 {
		return utils.NotFound("user not found.")
	}
	return nil
}

func Update(id, email, picture, fullName, givenName, familyName, description string) *utils.RestErr {
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
		return utils.InternalServerErr(err.Error())
	}
	if result.MatchedCount == 0 {
		return utils.NotFound("user not found.")
	}
	if result.ModifiedCount == 0 {
		return utils.BadRequest("invalid update request.")
	}
	return nil
}
