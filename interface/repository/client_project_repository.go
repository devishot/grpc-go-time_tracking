package repository

import (
  "github.com/devishot/grpc-go-time_tracking/app"
  "github.com/devishot/grpc-go-time_tracking/infrastructure/database"
  "github.com/devishot/grpc-go-time_tracking/infrastructure/database/table"
  "github.com/devishot/grpc-go-time_tracking/interface/factory"
)

type ProjectRepository struct {
  table *table.ClientProjectTable
}

func NewProjectRepository(db *database.DB) (*ProjectRepository, error) {
  t, err := table.NewClientProjectTable(db)
  if err != nil {
    return nil, err
  }

  return &ProjectRepository{table: t}, nil
}

func (tr *ProjectRepository) GetByID(id string) (*app.ProjectEntity, error) {
  f := factory.ProjectRowFactory{}

  row, err := tr.table.FindByID(id)
  if err != nil {
    return nil, err
  }

  return f.GetDomain(row), nil
}
