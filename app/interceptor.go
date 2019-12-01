package app

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	scopeKey = "scopes"
)

func AuthInterceptor(policy Policy) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		scopes, err := GetScopes(ctx)
		if err != nil {
			return nil, err
		}

		if !VerifyPermissions(scopes, policy[info.FullMethod]) {
			return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
		}
		return handler(ctx, req)
	}
}

// NOTE: dummy implementation of verification and retrieving user permission
func GetScopes(ctx context.Context) ([]string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	vs, ok := md[scopeKey]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, codes.Unauthenticated.String())
	}
	return vs, nil
}

func VerifyPermissions(having, required Permissions) bool {
	for _, r := range required {
		if !contains(having, r) {
			return false
		}
	}
	return true
}

func contains(s []string, target string) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
