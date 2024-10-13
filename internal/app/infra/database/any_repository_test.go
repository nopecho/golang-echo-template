package database

import (
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/util/gorm/datasource"
	"github.com/nopecho/golang-template/test/testcontainer"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	container = testcontainer.NewPostgresContainer()
	db        = datasource.NewPostgresWith(container.DSN, nil)
)

func TestDomainRepository(t *testing.T) {
	_ = db.AutoMigrate(&AnyEntity{})
	sut := NewAnyGormRepository(db)

	t.Run("FindAll", func(t *testing.T) {
		_, _ = sut.Save(domain.NewAnyModel("1"))
		_, _ = sut.Save(domain.NewAnyModel("2"))
		_, _ = sut.Save(domain.NewAnyModel("3"))

		actual, _ := sut.FindAll()

		assert.Len(t, actual, 3)
	})

	t.Cleanup(func() {
		container.Terminate()
	})
}
