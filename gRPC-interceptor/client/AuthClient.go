package main

import (
	"context"
	"fmt"

	authpb "github.com/sevendollar/grpcintercept/proto/auth"
	"google.golang.org/grpc"
)

type AuthClient struct {
	service  authpb.AuthServiceClient
	email    string
	password string
}

func newAuthClient(cc grpc.ClientConnInterface, email, password string) *AuthClient {
	return &AuthClient{
		service:  authpb.NewAuthServiceClient(cc),
		email:    email,
		password: password,
	}
}

func (auth *AuthClient) login() (token string, err error) {
	ctx := context.Background()

	resp, err := auth.service.Login(ctx, &authpb.LoginRequest{
		Email:    auth.email,
		Password: auth.password,
	})
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	return resp.GetAccessToken(), nil
}
