package factory

import (
	"github.com/devishot/grpc-go-time_tracking/app"
	"github.com/devishot/grpc-go-time_tracking/interface/api"
)

type TimeRecordDomainFactory struct {
	Message      *api.TimeRecord
	DomainObject *app.TimeRecordEntity
}

func NewTimeRecordDomainFactory(m *api.TimeRecord) *TimeRecordDomainFactory {
	f := &TimeRecordDomainFactory{Message: m}
	f.Build()
	return f
}

func (f *TimeRecordDomainFactory) Build() {
	f.DomainObject = &app.TimeRecordEntity{
		Id:          f.Message.Id,
		Amount:      f.Message.Amount,
		Timestamp:   f.Message.Timestamp,
		Description: f.Message.Description,
	}
}

func (f *TimeRecordDomainFactory) GetOwnerId() string {
	return f.Message.UserId
}

func (f *TimeRecordDomainFactory) GetProjectId() string {
	return f.Message.ProjectId
}
