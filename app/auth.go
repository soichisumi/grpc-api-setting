package app

import (
	"fmt"
	"github.com/jhump/protoreflect/grpcreflect"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

const (
	DATA_GET    = "data.get"
	DATA_SET    = "data.set"
	DATA_DELETE = "data.delete"
)

var (
	APIPolicy = Policy(map[string]Permissions{
		"/proto.Server/GetData":    NewPermissions(DATA_GET),
		"/proto.Server/SetData":    NewPermissions(DATA_SET),
		"/proto.Server/ListData":   NewPermissions(DATA_GET),
		"/proto.Server/DeleteData": NewPermissions(DATA_DELETE),
	})
)

// Permissions ...
type Permissions []string

func NewPermissions(scopes ...string) Permissions {
	return Permissions(scopes)
}

type Policy map[string]Permissions

// ref: https://github.com/grpc/grpc-go/issues/1526#issuecomment-372481817
func GetMethodDescriptor(s *grpc.Server) ([]string, error) {
	sds, err := grpcreflect.LoadServiceDescriptors(s)
	if err != nil {
		return nil, err
	}

	fullMethods := make([]string, 0)
	for _, sd := range sds {
		for _, md := range sd.GetMethods() {
			methodName := fmt.Sprintf("/%s/%s", sd.GetFullyQualifiedName(), md.GetName())
			fullMethods = append(fullMethods, methodName)
		}
	}
	return fullMethods, nil
}

func VerifyPolicy(fullMethods []string, policy Policy) error {
	fmt.Printf("debug: fullmethods: %v\n", fullMethods)
	// check policy whether all endpoints are set
	for _, v := range fullMethods {
		_, ok := policy[v]
		if !ok {
			return xerrors.New("policy for method is not set")
		}
	}
	return nil
}
