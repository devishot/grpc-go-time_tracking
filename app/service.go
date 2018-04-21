package app

type Service struct {
	TimeRecordRepository TimeRecordRepository
	UserRepository       UserRepository
	ProjectRepository    ProjectRepository
}

func (s *Service) CreateRecord(userID string, projectID string, record *TimeRecordEntity) (*TimeRecordEntity, error) {
	owner, err := s.UserRepository.GetById(userID)
	if err != nil {
		return nil, err
	}

	project, err := s.ProjectRepository.GetById(projectID)
	if err != nil {
		return nil, err
	}

	record.Owner = owner
	record.Project = project

	return s.TimeRecordRepository.Store(record)
}

func (s *Service) DeleteRecord(recordID string) error {
	return s.TimeRecordRepository.DeleteById(recordID)
}

func (s *Service) AllRecords(userID string, projectID string) ([]*TimeRecordEntity, error) {
	var records []*TimeRecordEntity
	var err error

	if userID != "" {
		records, err = s.TimeRecordRepository.GetByOwnerId(userID)
	} else if projectID != "" {
		records, err = s.TimeRecordRepository.GetByProjectId(projectID)
	}

	return records, err
}
