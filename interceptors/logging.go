package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func UnaryServerLoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("\n\t======== Logging Interceptor ========\nRPC Call: %s\nRequest: %v", info.FullMethod, req)
	res, err := handler(ctx, req)
	if err != nil {
		log.Printf("\nRPC Call: %s\nError: %v\n\t======== End Logging Interceptor ========", info.FullMethod, err)
	} else {
		log.Printf("\nRPC Call: %s\n\t======== End Logging Interceptor ========", info.FullMethod)
	}
	return res, err
}

func UnaryClientLoggingInterceptor(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	log.Printf("\n\t====== RPC Call Interceptor ======\nMethod: %s", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	log.Printf("\nMethod: %s\n\t====== End RPC Call Interceptor ======", method)
	return err
}
