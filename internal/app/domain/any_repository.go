package domain

type AnyRepository interface {
	FindAll() ([]*AnyModel, error)
	FindById(id uint64) (*AnyModel, error)
	Save(domain *AnyModel) (*AnyModel, error)
}
