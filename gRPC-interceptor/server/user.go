package main

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	UUID           uuid.UUID
	Email          string
	HashedPassword string
	Roles          []string
}

// Create the user
func newUser(email, password string) (*User, error) {
	email = strings.TrimSpace(email)
	password = strings.TrimSpace(password)

	byteHashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &User{
		UUID:           uuid.New(),
		Email:          email,
		HashedPassword: string(byteHashedPassword),
	}, nil
}

// Check the password
func (user *User) isCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

// Clone User
func (user *User) clone() *User {
	return &User{
		UUID:           user.UUID,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Roles:          user.Roles,
	}
}
