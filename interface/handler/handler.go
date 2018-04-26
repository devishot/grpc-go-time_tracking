package handler

import (
	"context"

	"github.com/devishot/grpc-go-time_tracking/app"
	"github.com/devishot/grpc-go-time_tracking/interface/api"
	"github.com/devishot/grpc-go-time_tracking/interface/factory"
)

// Server represents the gRPC server
type Server struct {
	TimeRecordRepository app.TimeRecordRepository
	ProjectRepository app.ProjectRepository
	UserRepository app.UserRepository
}

func (s *Server) CreateRecord(ctx context.Context, in *api.TimeRecord) (*api.TimeRecord, error) {
	f := factory.NewTimeRecordDomainFactory(in)

	record, err := s.AppService().CreateRecord(f.GetOwnerID(), f.GetProjectID(), f.DomainObject)
	if err != nil {
		return nil, err
	}

	return factory.NewTimeRecordMessageFactory(record).Message, nil
}

func (s *Server) DeleteRecord(ctx context.Context, in *api.DeleteRecordRequest) (*api.TimeRecord, error) {
	err := s.AppService().DeleteRecord(in.Id)
	msg := &api.TimeRecord{Id: in.Id}
	return msg, err
}

func (s *Server) AllRecords(ctx context.Context, in *api.AllRecordsRequest) (*api.TimeRecords, error) {
	records, err := s.AppService().AllRecords(in.GetUserId(), in.GetProjectId())
	if err != nil {
		return nil, err
	}

	return factory.NewTimeRecordsMessageFactory(records).Message, nil
}

func (s *Server) AppService() *app.Service {
	return &app.Service{
	  TimeRecordRepository: s.TimeRecordRepository,
    ProjectRepository: s.ProjectRepository,
    UserRepository: s.UserRepository,
  }
}
