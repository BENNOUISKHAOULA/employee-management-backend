package models

type Employee struct {
	ID         string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName  string `json:"firstName" bson:"firstName"`
	LastName   string `json:"lastName" bson:"lastName"`
	Email      string `json:"email" bson:"email"`
	Phone      string `json:"phone" bson:"phone"`
	Position   string `json:"position" bson:"position"`
	Department string `json:"department" bson:"department"`
	DateOfHire string `json:"dateOfHire" bson:"dateOfHire"`
}
