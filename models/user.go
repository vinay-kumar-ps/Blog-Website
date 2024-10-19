package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"` // Corrected the typo here
	Email     string `json:"email"`
	Password  []byte `json:"-"`         // Do not expose the password in JSON responses
	Phone     string `json:"phone"`
}

// SetPassword hashes the password and assigns it to the User model
func (user *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}
