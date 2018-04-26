package app

type Service struct {
	TimeRecordRepository TimeRecordRepository
	UserRepository       UserRepository
	ProjectRepository    ProjectRepository
}

func (s *Service) CreateRecord(userID string, projectID string, record *TimeRecordEntity) (*TimeRecordEntity, error) {
	/*TODO uncomment it
	owner, err := s.UserRepository.GetByID(userID)
	if err != nil {
		return nil, err
	}

	project, err := s.ProjectRepository.GetByID(projectID)
	if err != nil {
		return nil, err
	}

	record.Owner = owner
	record.Project = project
	*/
	record.generateID()
	record.generateTimestamp()

	return s.TimeRecordRepository.Store(record)
}

func (s *Service) DeleteRecord(recordID string) error {
	return s.TimeRecordRepository.DeleteByID(recordID)
}

func (s *Service) AllRecords(userID string, projectID string) ([]*TimeRecordEntity, error) {
	var records []*TimeRecordEntity
	var err error

	if userID != "" {
		records, err = s.TimeRecordRepository.GetByOwnerID(userID)
	} else if projectID != "" {
		records, err = s.TimeRecordRepository.GetByProjectID(projectID)
	}

	return records, err
}
