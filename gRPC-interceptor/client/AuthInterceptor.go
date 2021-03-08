package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type AuthInterceptor struct {
	authClient   *AuthClient
	acccessToken string
}

func newAuthInterceptor(authClient *AuthClient) *AuthInterceptor {
	return &AuthInterceptor{
		authClient: authClient,
	}
}

func (interceptor *AuthInterceptor) Unary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		// Logging
		// log.Println("--> Unary interceptor")

		// Get access token
		err := interceptor.refreshToken()
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		// Attach access token
		ctx = interceptor.attachAccessToken(ctx)

		return invoker(ctx, method, req, reply, cc, opts...)
	}

}

func (interceptor *AuthInterceptor) refreshToken() error {
	token, err := interceptor.authClient.login()
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	interceptor.acccessToken = token

	return nil
}

func (interceptor *AuthInterceptor) attachAccessToken(ctx context.Context) context.Context {
	return metadata.AppendToOutgoingContext(ctx, "authorization", interceptor.acccessToken)
}
