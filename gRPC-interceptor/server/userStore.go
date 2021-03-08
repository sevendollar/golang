package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-redis/redis/v8"
)

// UserStore interface
type UserStore interface {
	save(user *User) error
	find(email string) (*User, error)
}

// Implement memory store
type memoryStore struct {
	users map[string]*User
}

// Create a new memory store
func newMemoryStore(user *User) *memoryStore {
	return &memoryStore{
		users: make(map[string]*User),
	}
}

// Save to memory
func (store *memoryStore) save(user *User) error {
	_, ok := store.users[user.Email]
	if ok {
		return fmt.Errorf("user existed")
	}

	// Store to the memory
	store.users[user.Email] = user.clone()

	return nil
}

// Find from memory
func (store *memoryStore) find(email string) (*User, error) {
	email = strings.TrimSpace(email)
	user, ok := store.users[email]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// Implement Redis store
type redisStore struct {
	client *redis.Client
}

// Create a Redis client
func newRedisStore(addr, username, password string) *redisStore {
	addr = strings.TrimSpace(addr)
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})

	return &redisStore{
		client: c,
	}
}

func (store *redisStore) close() error {
	return store.client.Close()
}

// Save to Redis
func (store *redisStore) save(user *User) error {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	err = store.client.Set(context.Background(), user.Email, jsonUser, 0).Err()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Find from Redis
func (store *redisStore) find(email string) (*User, error) {
	email = strings.TrimSpace(email)

	jsonUser, err := store.client.Get(context.Background(), email).Result()
	if err != nil {
		return nil, fmt.Errorf("user do not exist: %w", err)
	}

	user := &User{}

	err = json.Unmarshal([]byte(jsonUser), user)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return user, nil
}

// Adds role to the User struct
func (store *redisStore) addRole(email string, role string) error {
	email = strings.TrimSpace(email)
	role = strings.Title(strings.ToLower(strings.TrimSpace(role)))

	user, err := store.find(email)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	for _, r := range user.Roles {
		if role == r {
			return fmt.Errorf("role has existed")
		}
	}

	user.Roles = append(user.Roles, role)
	err = store.save(user)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Deletes role from the User struct
func (store *redisStore) delRole(email string, removeRole string) error {
	email = strings.TrimSpace(email)
	removeRole = strings.Title(strings.ToLower(strings.TrimSpace(removeRole)))

	user, err := store.find(email)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	newRoles := []string{}
	roles := user.Roles

	for i, role := range roles {
		// Find matched role
		if role == removeRole {
			newRoles = append(newRoles, roles[:i]...)
			newRoles = append(newRoles, roles[i+1:]...)

			user.Roles = newRoles

			err = store.save(user)
			if err != nil {
				return fmt.Errorf("%w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("role not found")
}
