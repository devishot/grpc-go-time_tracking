package handler

import (
	"context"

	"github.com/devishot/grpc-go-time_tracking/api"
)

// Server represents the gRPC server
type Server struct {
}

func (s *Server) CreateRecord(ctx context.Context, in *api.TimeRecord) (*api.TimeRecord, error) {
	return in, nil
}

func (s *Server) DeleteRecord(ctx context.Context, in *api.DeleteRecordRequest) (*api.TimeRecord, error) {
	out := &api.TimeRecord{}
	out.Id = in.Id
	return out, nil
}

func (s *Server) AllRecords(ctx context.Context, in *api.AllRecordsRequest) (*api.TimeRecords, error) {
	results := &api.TimeRecords{}
	return results, nil
}
