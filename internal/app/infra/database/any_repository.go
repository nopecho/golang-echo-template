package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"gorm.io/gorm"
)

type AnyGormRepository struct {
	db *gorm.DB
}

func NewAnyGormRepository(db *gorm.DB) *AnyGormRepository {
	return &AnyGormRepository{
		db: db,
	}
}

func (r *AnyGormRepository) FindAll() ([]*domain.AnyModel, error) {
	var entities []AnyEntity
	if err := r.db.Find(&entities).Error; err != nil {
		return []*domain.AnyModel{}, err
	}

	var domains []*domain.AnyModel
	for _, e := range entities {
		domains = append(domains, e.Domain())
	}
	return domains, nil
}

func (r *AnyGormRepository) FindById(id uint64) (*domain.AnyModel, error) {
	var entity AnyEntity
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return entity.Domain(), nil
}

func (r *AnyGormRepository) Save(domain *domain.AnyModel) (*domain.AnyModel, error) {
	entity := AnyEntity{
		BaseModel: BaseModel{
			ID: domain.ID,
		},
	}
	if err := r.db.Create(&entity).Error; err != nil {
		return nil, err
	}
	return entity.Domain(), nil
}
