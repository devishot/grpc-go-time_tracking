package factory

import (
	"github.com/devishot/grpc-go-time_tracking/app"
	"github.com/devishot/grpc-go-time_tracking/infrastructure/database/table"
)

type TimeRecordRowFactory struct {
}

func (f *TimeRecordRowFactory) GetRow(d *app.TimeRecordEntity) table.TimeRecordRow {
	return table.TimeRecordRow{
		ID:          d.ID,
		Amount:      d.Amount,
		Timestamp:   d.Timestamp,
		Description: d.Description,
		UserID:      d.GetUserID(),
		ProjectID:   d.GetProjectID(),
	}
}

func (f *TimeRecordRowFactory) GetDomain(r table.TimeRecordRow) *app.TimeRecordEntity {
	d := &app.TimeRecordEntity{
		ID:          r.ID,
		Amount:      r.Amount,
		Timestamp:   r.Timestamp,
		Description: r.Description,
	}

	if uid := r.UserID; uid != "" {
		d.Owner = &app.UserEntity{ID: uid}
	}

	if pid := r.ProjectID; pid != "" {
		d.Project = &app.ProjectEntity{ID: pid}
	}

	return d
}


type UserRowFactory struct {
}

func (f *UserRowFactory) GetRow(d *app.UserEntity) table.UserRow {
  return table.UserRow{
    ID:          d.ID,
  }
}

func (f *UserRowFactory) GetDomain(r table.UserRow) *app.UserEntity {
  d := &app.UserEntity{
    ID:          r.ID,
  }

  return d
}


type ProjectRowFactory struct {
}

func (f *ProjectRowFactory) GetRow(d *app.ProjectEntity) table.ClientProjectRow {
  return table.ClientProjectRow{
    ID:          d.ID,
  }
}

func (f *ProjectRowFactory) GetDomain(r table.ClientProjectRow) *app.ProjectEntity {
  d := &app.ProjectEntity{
    ID:          r.ID,
  }

  return d
}