package models

type Employee struct {
    ID         string `json:"id,omitempty"`
    FirstName  string `json:"firstName"`
    LastName   string `json:"lastName"`
    Email      string `json:"email"`
    Phone      string `json:"phone"`
    Position   string `json:"position"`
    Department string `json:"department"`
    DateOfHire string `json:"dateOfHire"`
}
