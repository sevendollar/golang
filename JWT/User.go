package main

import (
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string
	HashedPassword string
	Role           string
}

func newUser(username, password string) (*User, error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &User{
		Username:       username,
		HashedPassword: string(hashedPassword),
	}, nil
}

func (u *User) isRightPassword(password string) bool {
	password = strings.TrimSpace(password)

	err := bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password))
	return err == nil
}
