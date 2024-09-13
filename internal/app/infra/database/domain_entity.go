package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"strconv"
)

type DomainEntity struct {
	BaseModel
}

func (DomainEntity) TableName() string {
	return "domain"
}

func (e *DomainEntity) Domain() *domain.Domain {
	return &domain.Domain{
		ID:   e.ID,
		Name: strconv.FormatUint(e.ID, 10),
	}
}
