// This interceptor will contain a JWT manager and a map that define for each RPC method a list of roles that can access it. The key of the map is the full method name, and its value is a slice of role names.

package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	// metadata
	metadataAuthorize = "authorization"
)

func accessbleRoles() map[string][]string {
	// authServicePath := "/github.com.sevendollar.grpcintercept.AuthService/"
	helloServicePath := "/github.com.sevendollar.grpcintercept.HelloService/"
	worldServicePath := "/github.com.sevendollar.grpcintercept.WorldService/"

	return map[string][]string{
		// No roles mean block everything
		helloServicePath + "Man": {},
		helloServicePath + "Cat": {"Admin"},

		worldServicePath + "Tuna": {"PowerUser", "Admin"},
	}
}

type AuthInterceptor struct {
	accessbleRoles map[string][]string
	jwtManager     *JWTManager
}

func newAuthInterceptor(accessbleRoles map[string][]string, jwtManager *JWTManager) *AuthInterceptor {
	return &AuthInterceptor{
		accessbleRoles: accessbleRoles,
		jwtManager:     jwtManager,
	}
}

// Unary Auth interceptor for unary requests
func (interceptor *AuthInterceptor) Unary() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {

		// Logging
		log.Println("--> Unary interceptor:", info.FullMethod)

		// Show metadata
		md, _ := metadata.FromIncomingContext(ctx)
		fmt.Println("metadata:", md)

		// Authorization
		err = interceptor.authorize(ctx, info)
		if err != nil {
			return nil, fmt.Errorf("%w", err)
		}

		return handler(ctx, req)
	}
}

func (interceptor *AuthInterceptor) authorize(ctx context.Context, info *grpc.UnaryServerInfo) error {
	acls := accessbleRoles()
	roles, ok := acls[info.FullMethod]
	if !ok {
		return nil
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return fmt.Errorf("metadata not found")
	}
	values, ok := md[metadataAuthorize]
	if !ok {
		return fmt.Errorf("token not found")

	}
	token := values[0]
	claims, err := interceptor.jwtManager.verify(token)
	if err != nil {
		return fmt.Errorf("invalid token: %w", err)
	}

	for _, accessRole := range roles {
		for _, userRole := range claims.Roles {
			if userRole == accessRole {
				return nil
			}
		}
	}

	return fmt.Errorf("you do not have any permission")
}
