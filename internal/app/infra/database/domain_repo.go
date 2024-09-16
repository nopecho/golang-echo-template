package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"gorm.io/gorm"
)

type DomainPostgresRepository struct {
	db *gorm.DB
}

func NewDomainPostgresRepository(db *gorm.DB) *DomainPostgresRepository {
	return &DomainPostgresRepository{
		db: db,
	}
}

func (r *DomainPostgresRepository) FindAll() ([]*domain.Domain, error) {
	var entities []DomainEntity
	if err := r.db.Find(&entities).Error; err != nil {
		return []*domain.Domain{}, err
	}

	var domains []*domain.Domain
	for _, e := range entities {
		domains = append(domains, e.Domain())
	}
	return domains, nil
}

func (r *DomainPostgresRepository) FindById(id uint64) (*domain.Domain, error) {
	var entity DomainEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return entity.Domain(), nil
}

func (r *DomainPostgresRepository) Save(domain *domain.Domain) (*domain.Domain, error) {
	entity := DomainEntity{
		BaseModel: BaseModel{
			ID: domain.ID,
		},
	}
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return entity.Domain(), nil
}
