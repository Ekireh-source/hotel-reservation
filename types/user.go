package types

import (
	"fmt"
	"regexp"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)


var(
	bcryptCost = 10 // bcrypt cost factor
	minFirstNameLength = 2 // minimum length for first name
	minLastNameLength = 2 // minimum length for last name
	minPasswordLength = 8 // minimum length for password

)

type CreateUserParams struct {
	FirstName string `json:"firstName" `
	LastName  string `json:"lastName" `
	Email     string `json:"email" `
	Password  string `json:"password"`
}

func (p *CreateUserParams) Validate() []string {

	errors := []string{}
	if len(p.FirstName) < minFirstNameLength {
		errors = append(errors, fmt.Sprintf("first name must be at least %d characters long", minFirstNameLength))
	}
	if len(p.LastName) < minLastNameLength {
		errors = append(errors, fmt.Sprintf("last name must be at least %d characters long", minLastNameLength))
	}
	
	if len(p.Password) < minPasswordLength {
		errors = append(errors, fmt.Sprintf("password must be at least %d characters long", minPasswordLength))
	}
	if !isEmailValid(p.Email) {
		errors = append(errors, "email is not valid")
	}
	return errors
}

func isEmailValid(email string) bool {
	// A simple email validation logic, can be improved with regex
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	
	return emailRegex.MatchString(email)
}

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
	Email	 string `bson:"email" json:"email"`
	EncryptedPassword string `bson:"encryptedPassword" json:"-"`
	
}

func NewUserFromParams(params CreateUserParams) (*User, error) {
	encpw, err :=  bcrypt.GenerateFromPassword([]byte(params.Password), bcryptCost)
	if err != nil {
		return nil, err
	}
	return &User{
		FirstName: params.FirstName,
		LastName:  params.LastName,
		Email:     params.Email,
		EncryptedPassword: string(encpw),
	}, nil
}
