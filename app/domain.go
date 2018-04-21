package app

import (
	"time"

	"github.com/satori/go.uuid"
)

type TimeRecordRepository interface {
	Store(*TimeRecordEntity) (*TimeRecordEntity, error)
	DeleteByID(string) error
	GetByID(string) (*TimeRecordEntity, error)
	GetByOwnerID(string) ([]*TimeRecordEntity, error)
	GetByProjectID(string) ([]*TimeRecordEntity, error)
}

type UserRepository interface {
	GetByID(string) (*UserEntity, error)
}

type ProjectRepository interface {
	GetByID(string) (*ProjectEntity, error)
}

type TimeRecordEntity struct {
	ID          string
	Amount      int32
	Timestamp   time.Time
	Description string
	Owner       *UserEntity
	Project     *ProjectEntity
}

func (t *TimeRecordEntity) GetUserID() string {
	owner := t.Owner
	if owner != nil {
		return owner.ID
	}
	return ""
}

func (t *TimeRecordEntity) GetProjectID() string {
	project := t.Project
	if project != nil {
		return project.ID
	}
	return ""
}

func (t *TimeRecordEntity) generateID() {
	uuid := uuid.Must(uuid.NewV4())
	t.ID = uuid.String()
}

func (t *TimeRecordEntity) generateTimestamp() {
	t.Timestamp = time.Now()
}

type UserEntity struct {
	ID string
}

type ProjectEntity struct {
	ID string
}
