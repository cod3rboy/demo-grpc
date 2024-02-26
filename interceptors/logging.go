package interceptors

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

func LoggingInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	log.Printf("\n\t======== Logging Interceptor ========\nRPC Call: %s\nRequest: %v", info.FullMethod, req)
	res, err := handler(ctx, req)
	if err != nil {
		log.Printf("\nRPC Call: %s\nError: %v\n\t======== End Logging Interceptor ========", info.FullMethod, err)
	} else {
		log.Printf("\nRPC Call: %s\n\t======== End Logging Interceptor ========", info.FullMethod)
	}
	return res, err
}
