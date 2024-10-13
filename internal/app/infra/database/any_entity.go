package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"strconv"
)

type AnyEntity struct {
	BaseModel
}

func (AnyEntity) TableName() string {
	return "domain"
}

func (e *AnyEntity) Domain() *domain.AnyModel {
	return &domain.AnyModel{
		ID:   e.ID,
		Name: strconv.FormatUint(uint64(e.ID), 10),
	}
}
