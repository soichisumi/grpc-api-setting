package app

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/soichisumi/grpc-api-setting/app/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServer() proto.ServerServer {
	return &Server{
		db: make(map[string]string),
	}
}

type Server struct {
	db map[string]string //NOTE: map is not safe for concurrent use. should use sync.Map
}

func (s *Server) GetData(ctx context.Context, req *proto.GetDataRequest) (*proto.GetDataResponse, error) {
	v, ok := s.db[req.Key]
	if !ok {
		return nil, status.Error(codes.NotFound, codes.NotFound.String())
	}
	return &proto.GetDataResponse{
		Value: v,
	}, nil
}

func (s *Server) SetData(ctx context.Context, req *proto.SetDataRequest) (*empty.Empty, error) {
	s.db[req.Key] = req.Value
	return &empty.Empty{}, nil
}

func (s *Server) ListData(ctx context.Context, req *empty.Empty) (*proto.ListDataResponse, error) {
	res := make([]*proto.Data, 0, len(s.db))
	for k, v := range s.db {
		res = append(res, &proto.Data{
			Key:   k,
			Value: v,
		})
	}
	return &proto.ListDataResponse{
		Data: res,
	}, nil
}

func (s *Server) DeleteData(ctx context.Context, req *proto.DeleteDataRequest) (*empty.Empty, error) {
	delete(s.db, req.Key)
	return &empty.Empty{}, nil
}
