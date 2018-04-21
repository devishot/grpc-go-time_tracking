package app

type TimeRecordRepository interface {
	Store(*TimeRecordEntity) (*TimeRecordEntity, error)
	DeleteById(string) error
	GetById(string) (*TimeRecordEntity, error)
	GetByOwnerId(string) ([]*TimeRecordEntity, error)
	GetByProjectId(string) ([]*TimeRecordEntity, error)
}

type UserRepository interface {
	GetById(string) (*UserEntity, error)
}

type ProjectRepository interface {
	GetById(string) (*ProjectEntity, error)
}

type TimeRecordEntity struct {
	Id          string
	Amount      int32
	Timestamp   int64
	Description string
	Owner       *UserEntity
	Project     *ProjectEntity
}

type UserEntity struct {
	Id string
}

type ProjectEntity struct {
	Id string
}
