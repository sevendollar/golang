package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func helloInterceptorUnary() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		ctx = metadata.AppendToOutgoingContext(ctx,
			"helloInterceptorUnary", "foo",
			"helloInterceptorUnary", "bar",
		)

		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
