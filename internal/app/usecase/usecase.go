package usecase

import (
	"github.com/nopecho/golang-template/internal/app/domain/repository"
	"github.com/nopecho/golang-template/internal/app/infra/clinet"
	"github.com/nopecho/golang-template/internal/app/infra/db"
)

type Loader interface {
	Load() (interface{}, error)
}

type DomainUsecase struct {
	Repository repository.DomainRepository
	Client     *clinet.Client
}

func NewDomainUsecase(client *clinet.Client) *DomainUsecase {
	return &DomainUsecase{
		Repository: db.NewDomainPostgresRepository(),
		Client:     client,
	}
}

func (u *DomainUsecase) Load() (interface{}, error) {
	return map[string]any{
		"key": "value",
	}, nil
}
