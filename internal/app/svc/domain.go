package svc

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/app/infra/clinet"
	"github.com/nopecho/golang-template/internal/app/svc/cmd"
)

type DomainService struct {
	repository domain.Repository
	client     *clinet.Client
}

func NewDomainService(repository domain.Repository, client *clinet.Client) *DomainService {
	return &DomainService{
		repository: repository,
		client:     client,
	}
}

func (s *DomainService) GetById(id uint64) (*domain.Domain, error) {
	return s.repository.FindById(id)
}

func (s *DomainService) Create(cmd cmd.DomainCreateCommand) (*domain.Domain, error) {
	d := domain.NewDomain(cmd.Name)
	created, err := s.repository.Save(d)
	if err != nil {
		return nil, err
	}
	return created, nil
}

func (s *DomainService) Update(cmd cmd.DomainUpdateCommand) (*domain.Domain, error) {
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
