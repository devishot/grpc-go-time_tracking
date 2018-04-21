package factory

import (
	"github.com/devishot/grpc-go-time_tracking/app"
	"github.com/devishot/grpc-go-time_tracking/interface/api"
)

type TimeRecordMessageFactory struct {
	DomainObject *app.TimeRecordEntity
	Message      *api.TimeRecord
}

func NewTimeRecordMessageFactory(obj *app.TimeRecordEntity) *TimeRecordMessageFactory {
	f := &TimeRecordMessageFactory{DomainObject: obj}
	f.Build()
	return f
}

func (f *TimeRecordMessageFactory) Build() {
	f.Message = &api.TimeRecord{
		Id:          f.DomainObject.ID,
		Amount:      f.DomainObject.Amount,
		Timestamp:   f.DomainObject.Timestamp.Unix(),
		Description: f.DomainObject.Description,
		UserId:      f.GetUserID(),
		ProjectId:   f.GetProjectID(),
	}
}

func (f *TimeRecordMessageFactory) GetUserID() string {
	return f.DomainObject.GetUserID()
}

func (f *TimeRecordMessageFactory) GetProjectID() string {
	return f.DomainObject.GetProjectID()
}

type TimeRecordsMessageFactory struct {
	Objects []*app.TimeRecordEntity
	Message *api.TimeRecords
}

func NewTimeRecordsMessageFactory(objs []*app.TimeRecordEntity) *TimeRecordsMessageFactory {
	f := &TimeRecordsMessageFactory{Objects: objs}
	f.Build()
	return f
}

func (f *TimeRecordsMessageFactory) Build() {
	records := make([]*api.TimeRecord, len(f.Objects))
	for i, v := range f.Objects {
		records[i] = NewTimeRecordMessageFactory(v).Message
	}
	f.Message = &api.TimeRecords{Records: records}
}
