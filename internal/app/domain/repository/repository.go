package repository

import "github.com/nopecho/golang-template/internal/app/domain"

type DomainRepository interface {
	FindAll() ([]domain.Domain, error)
	FindById(id int) (domain.Domain, error)
	Create(domain domain.Domain) error
}
