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
	// TODO: Add another user.
	// TODO: Try Create func on a user that exists.
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
	// TODO: Try to retrive a user that doesn't exists and check for expected errors.
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
	// TODO: Try the update func with the same values and check for errors.
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

func TestAddGroup(t *testing.T) {
	// TODO: Try to add a group that already is a member and check for errors.
	restErr := AddGroup("1", "A")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
	restErr = AddGroup("1", "B")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
}

func TestDelGroup(t *testing.T) {
	// TODO: Try to delete a group that the user is not a member of.
	restErr := DelGroup("1", "A")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
}

func TestAddContact(t *testing.T) {
	// TODO: Try to add a contact that already exists and check for errors.
	restErr := AddContact("1", "a")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
	restErr = AddContact("1", "b")
	if restErr != nil {
		t.Errorf(restErr.Message)
	}
}

func TestDelContact(t *testing.T) {
	// TODO: Try to delete a contact that doesn't exist.
	restErr := DelContact("1", "a")
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
