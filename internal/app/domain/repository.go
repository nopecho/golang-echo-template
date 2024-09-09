package domain

type Repository interface {
	FindAll() ([]*Domain, error)
	FindById(id int64) (*Domain, error)
	Save(domain *Domain) (*Domain, error)
}
