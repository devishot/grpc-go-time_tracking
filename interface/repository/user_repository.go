package repository

import (
  "github.com/devishot/grpc-go-time_tracking/app"
  "github.com/devishot/grpc-go-time_tracking/infrastructure/database"
  "github.com/devishot/grpc-go-time_tracking/infrastructure/database/table"
  "github.com/devishot/grpc-go-time_tracking/interface/factory"
)

type UserRepository struct {
  table *table.UserTable
}

func NewUserRepository(db *database.DB) (*UserRepository, error) {
  t, err := table.NewUserTable(db)
  if err != nil {
    return nil, err
  }

  return &UserRepository{table: t}, nil
}

func (tr *UserRepository) GetByID(id string) (*app.UserEntity, error) {
  f := factory.UserRowFactory{}

  row, err := tr.table.FindByID(id)
  if err != nil {
    return nil, err
  }

  return f.GetDomain(row), nil
}
