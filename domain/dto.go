package domain

// User type is the data transfer object that represent a user in the platform.
type User struct {
	// These fields are provided by google.
	ID         string `json:"id" bson:"_id"`
	Email      string `json:"email" bson:"email"`
	Picture    string `json:"picture" bson:"picture"`
	FullName   string `json:"fullName" bson:"fullName"`
	GivenName  string `json:"givenName" bson:"givenName"`
	FamilyName string `json:"familyName" bson:"familyName"`
	// These fields are platform specific.
	DateCreated    int64    `json:"dateCreated" bson:"dateCreated"` // Unix time.
	Description    string   `json:"description" bson:"description"`
	ContactIDs     []string `json:"contactIds" bson:"contactIds"`
	JoinedGroupIDs []string `json:"joinedGroupIds" bson:"joinedGroupIds"`
}
