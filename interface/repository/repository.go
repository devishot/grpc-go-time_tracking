package repository

import (
	"github.com/devishot/grpc-go-time_tracking/app"
	"github.com/devishot/grpc-go-time_tracking/infrastructure/database"
	"github.com/devishot/grpc-go-time_tracking/infrastructure/database/table"
	"github.com/devishot/grpc-go-time_tracking/interface/factory"
)

type TimeRecordRepository struct {
	table *table.TimeRecordTable
}

func NewTimeRecordRepository(db *database.DB) (*TimeRecordRepository, error) {
	table, err := table.NewTimeRecordTable(db)
	if err != nil {
		return nil, err
	}

	return &TimeRecordRepository{table: table}, nil
}

func (tr *TimeRecordRepository) Store(obj *app.TimeRecordEntity) (*app.TimeRecordEntity, error) {
	f := factory.TimeRecordRowFactory{}

	row, err := tr.table.Insert(f.GetRow(obj))
	if err != nil {
		return nil, err
	}

	return f.GetDomain(row), nil
}

func (tr *TimeRecordRepository) DeleteByID(id string) error {
	return nil
}

func (tr *TimeRecordRepository) GetByID(id string) (*app.TimeRecordEntity, error) {
	return &app.TimeRecordEntity{}, nil
}

func (tr *TimeRecordRepository) GetByOwnerID(userID string) ([]*app.TimeRecordEntity, error) {
	return make([]*app.TimeRecordEntity, 0), nil
}

func (tr *TimeRecordRepository) GetByProjectID(projectID string) ([]*app.TimeRecordEntity, error) {
	return make([]*app.TimeRecordEntity, 0), nil
}
