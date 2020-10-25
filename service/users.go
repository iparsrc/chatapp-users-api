package service

import (
	"time"

	"github.com/parsaakbari1209/ChatApp-users-api/domain"
	"github.com/parsaakbari1209/ChatApp-users-api/utils"
)

// Users interface defines the users service.
type Users interface {
	Create(newUser *domain.User) (*domain.User, *utils.RestErr)
	Retrive(id string, private bool) (*domain.User, *utils.RestErr)
	Delete(id string) *utils.RestErr
	Update(id, email, picture, fullName, givenName, familyName, description string) *utils.RestErr
	AddGroup(id, groupID string) *utils.RestErr
	DelGroup(id, groupID string) *utils.RestErr
	AddContact(id, contactID string) *utils.RestErr
	DelContact(id, contactID string) *utils.RestErr
}

type service struct {
	repo domain.Repository
}

// New returns a new user service.
func New() Users {
	return &service{
		repo: domain.NewRepository(),
	}
}

func (s *service) Create(newUser *domain.User) (*domain.User, *utils.RestErr) {
	// 1. Set the creation date.
	newUser.DateCreated = time.Now().UTC().Unix()

	// 2. Create the new user in the repository.
	user, restErr := s.repo.Create(newUser)

	// 3. Return values.
	return user, restErr
}

func (s *service) Retrive(id string, needPrivateProfile bool) (*domain.User, *utils.RestErr) {
	// 1. Retrive data from the repository.
	user, restErr := s.repo.Retrive(id)
	if restErr != nil {
		return nil, restErr
	}

	// 2. Check if the private profile is requested.
	if needPrivateProfile {
		return user, nil
	}

	// 3. Delete the user's private data.
	user.Email = ""
	user.FullName = ""
	user.FamilyName = ""
	user.ContactIDs = nil
	user.JoinedGroupIDs = nil

	// 4. Return values.
	return user, nil
}

func (s *service) Delete(id string) *utils.RestErr {
	// 1. Delete the user form the repository.
	restErr := s.repo.Delete(id)

	// 2. Return values.
	return restErr
}

func (s *service) Update(id, email, picture, fullName, givenName, familyName, description string) *utils.RestErr {
	// 1. Update the user in the repository.
	restErr := s.repo.Update(id, email, picture, fullName, givenName, familyName, description)

	// 2. Return values.
	return restErr
}

func (s *service) AddGroup(id, groupID string) *utils.RestErr {
	// 1. Add the group id to the user's profile.
	restErr := s.repo.AddGroup(id, groupID)

	// 2. Return values.
	return restErr
}

func (s *service) DelGroup(id, groupID string) *utils.RestErr {
	// 1. Remove the group id from the user's profile.
	restErr := s.repo.DelGroup(id, groupID)

	// 2. Return values.
	return restErr
}

func (s *service) AddContact(id, contactID string) *utils.RestErr {
	// 1. Add the contact id to the user's profile.
	restErr := s.repo.AddContact(id, contactID)

	// 2. Return values
	return restErr
}

func (s *service) DelContact(id, contactID string) *utils.RestErr {
	// 1. Remove the contact id from the user's profile.
	restErr := s.repo.DelContact(id, contactID)

	// 2. Return values.
	return restErr
}
