package domain

type User struct {
	// These fields are provided by google.
	ID         string `json:"id" bson:"_id"`
	Email      string `json:"email" bson:"email"`
	Picture    string `json:"picture" bson:"picture"`
	FullName   string `json:"fullName" bson:"fullName"`
	GivenName  string `json:"givenName" bson:"givenName"`
	FamilyName string `json:"familyName" bson:"familyName"`
	// These fields are platform specific.
	ContactsIDs    []string `json:"contactsIds" bson:"contactsIds"`
	JoinedGroupIDs []string `json:"joinedGroupIds" bson:"joinedGroupIds"`
}
