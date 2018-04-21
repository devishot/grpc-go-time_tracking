package factory

import (
	"time"

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
		ID:          f.Message.Id,
		Amount:      f.Message.Amount,
		Timestamp:   time.Unix(f.Message.Timestamp, 0),
		Description: f.Message.Description,
	}
}

func (f *TimeRecordDomainFactory) GetOwnerID() string {
	return f.Message.UserId
}

func (f *TimeRecordDomainFactory) GetProjectID() string {
	return f.Message.ProjectId
}
