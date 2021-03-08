package main

import (
	"context"
	"fmt"

	authpb "github.com/sevendollar/grpcintercept/proto/auth"
)

// AuthServer ...
type AuthServer struct {
	userStore  UserStore
	jwtManager *JWTManager
}

func newAuthServer(store UserStore, manager *JWTManager) *AuthServer {
	return &AuthServer{
		userStore:  store,
		jwtManager: manager,
	}
}

// Ping Implement gRPC AuthService/Ping method
func (auth *AuthServer) Ping(context.Context, *authpb.PingRequest) (*authpb.PingResponse, error) {
	return &authpb.PingResponse{
		Message: "Pong",
	}, nil
}

// Login Implement gRPC AuthService/Login method
func (auth *AuthServer) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginResponse, error) {
	user, err := auth.userStore.find(req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	accessToken, err := auth.jwtManager.generate(user)
	if err != nil {
		return nil, fmt.Errorf("%w", err)
	}

	return &authpb.LoginResponse{
		AccessToken: accessToken,
	}, nil
}
