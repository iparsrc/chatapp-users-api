package domain

import (
	"fmt"
	"testing"
)

const (
	uri = "mongodb://localhost:27017"
)

func TestConnectDB(t *testing.T) {
	ConnectDB(uri)
}

func TestCreate(t *testing.T) {
	user := User{
		ID:          "1",
		GivenName:   "Parsa",
		FamilyName:  "Akbari",
		Description: "I am 19.",
		FullName:    "Parsa Akbari",
		Email:       "akbariparsa1209@gmail.com",
	}
	result, restErr := Create(&user)
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
	if result == nil {
		t.Errorf("User is not created.")
	}
}

func TestRetrive(t *testing.T) {
	privateUser, restErr := Retrive("1", true)
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
	if privateUser == nil {
		t.Errorf("Private user with id 1 must not be nil.")
	}
	fmt.Println("   ", *privateUser)
	publicUser, restErr := Retrive("1", false)
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
	if publicUser == nil {
		t.Errorf("Public user with id 1 mush not be nil.")
	}
	fmt.Println("   ", *publicUser)
}

func TestUpdate(t *testing.T) {
	restErr := Update(
		"1",
		"parsaakbari80808080@gmail.com",
		"",
		"Parsa Akbari",
		"Parsa",
		"Akbari",
		"I am a devloper.",
	)
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
}

func TestDelete(t *testing.T) {
	restErr := Delete("1")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
}
