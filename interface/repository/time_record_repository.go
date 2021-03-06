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
	t, err := table.NewTimeRecordTable(db)
	if err != nil {
		return nil, err
	}

	return &TimeRecordRepository{table: t}, nil
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
	return tr.table.Delete(id)
}

func (tr *TimeRecordRepository) GetByID(id string) (*app.TimeRecordEntity, error) {
	f := factory.TimeRecordRowFactory{}

	row, err := tr.table.FindByID(id)
	if err != nil {
		return nil, err
	}

	return f.GetDomain(row), nil
}

func (tr *TimeRecordRepository) GetByOwnerID(userID string) (results []*app.TimeRecordEntity, err error) {
  f := factory.TimeRecordRowFactory{}

  rows, err := tr.table.FindByUserID(userID)
  if err != nil {
    return
  }

  for _, v := range rows {
    results = append(results, f.GetDomain(v))
  }

	return
}

func (tr *TimeRecordRepository) GetByProjectID(projectID string) (results []*app.TimeRecordEntity, err error) {
  f := factory.TimeRecordRowFactory{}

  rows, err := tr.table.FindByProjectID(projectID)
  if err != nil {
    return
  }

  for _, v := range rows {
    results = append(results, f.GetDomain(v))
  }

  return
}
