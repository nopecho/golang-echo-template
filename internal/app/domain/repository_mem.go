package domain

type MemoryRepository struct {
	Map map[int64]*Domain
}

func NewMemoryDomainRepository() *MemoryRepository {
	return &MemoryRepository{
		Map: make(map[int64]*Domain),
	}
}

func (r *MemoryRepository) FindAll() ([]*Domain, error) {
	var domains []*Domain
	for _, d := range r.Map {
		domains = append(domains, d)
	}
	return domains, nil
}

func (r *MemoryRepository) FindById(id int64) (*Domain, error) {
	return r.Map[id], nil
}

func (r *MemoryRepository) Save(domain *Domain) (*Domain, error) {
	r.Map[domain.ID] = domain
	return domain, nil
}
