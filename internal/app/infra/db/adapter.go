package db

import "github.com/nopecho/golang-template/internal/app/domain"

type DomainPostgresRepository struct {
	// db *sql.DB
}

func NewDomainPostgresRepository() *DomainPostgresRepository {
	return &DomainPostgresRepository{}
}

func (r *DomainPostgresRepository) FindAll() ([]domain.Domain, error) {
	return []domain.Domain{}, nil
}

func (r *DomainPostgresRepository) FindById(id int) (domain.Domain, error) {
	return domain.Domain{}, nil
}

func (r *DomainPostgresRepository) Create(domain domain.Domain) error {
	return nil
}
