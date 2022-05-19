package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func createUser(input RegisterInput) error {
	pwdHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return err
	}

	user := User{
		Email:     input.Email,
		Name:      input.Name,
		Password:  string(pwdHash),
		CreatedAt: time.Time{},
		Generated: 0,
	}
	collection.save(input.Email, &user)
	return nil
}

func authenticate(email, password string) error {
	user := collection.get(email)
	if user == nil {
		errors.New("no such user")
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
