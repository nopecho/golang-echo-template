package domain

type AnyService struct {
	repository AnyRepository
}

func NewAnyService(repository AnyRepository) *AnyService {
	return &AnyService{
		repository: repository,
	}
}

func (s *AnyService) GetById(id uint64) (*AnyModel, error) {
	return s.repository.FindById(id)
}

func (s *AnyService) Create(cmd *AnyCreateCommand) (*AnyModel, error) {
	d := NewAnyModel(cmd.Name)
	created, err := s.repository.Save(d)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *AnyService) Update(cmd *AnyUpdateCommand) (*AnyModel, error) {
	d, err := s.repository.FindById(cmd.ID)
	if err != nil || d == nil {
		return nil, err
	}
	d.Name = cmd.Name
	updated, err := s.repository.Save(d)
	if err != nil {
		return nil, err
	}
	return updated, nil
}
