package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/pkg/gorm/datasource"
	"github.com/nopecho/golang-template/test/testcontainer"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	container = testcontainer.NewPostgresContainer()
	db        = datasource.NewPostgres(container.DSN, nil)
)

func TestDomainRepository(t *testing.T) {
	_ = db.AutoMigrate(&DomainEntity{})
	sut := NewDomainPostgresRepository(db)

	t.Run("FindAll", func(t *testing.T) {
		_, _ = sut.Save(domain.NewDomain("1"))
		_, _ = sut.Save(domain.NewDomain("2"))
		_, _ = sut.Save(domain.NewDomain("3"))

		actual, _ := sut.FindAll()

		assert.Len(t, actual, 3)
	})

	t.Cleanup(func() {
		container.Terminate()
	})
}
