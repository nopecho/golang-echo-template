package domain

type Repository interface {
	FindAll() ([]*Domain, error)
	FindById(id uint64) (*Domain, error)
	Save(domain *Domain) (*Domain, error)
}
