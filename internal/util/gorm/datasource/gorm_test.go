package datasource

import (
	"context"
	"github.com/goccy/go-json"
	"github.com/nopecho/golang-template/internal/app/domain"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/test/testcontainer"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewPostgres(t *testing.T) {
	container := testcontainer.NewPostgresContainer()

	t.Run("TestNewPostgres", func(t *testing.T) {
		sut := NewPostgresWith(container.DSN, defaultConnPool())

		var actual []string
		sut.Raw("SELECT 1").Scan(&actual)

		assert.Len(t, actual, 1)
	})

	t.Run("AutoMigrate", func(t *testing.T) {
		sut := NewPostgresWith(container.DSN, defaultConnPool())
		_ = sut.AutoMigrate(&database.AnyEntity{})

		var actual []map[string]interface{}
		sut.Raw(`SELECT column_name, data_type
						FROM information_schema.columns
						WHERE table_name = 'domain'`).
			Scan(&actual)
		marshal, _ := json.Marshal(actual)
		t.Log(string(marshal))

		assert.Len(t, actual, 4)
	})

	t.Run("Transaction Manager Commit Test", func(t *testing.T) {
		source := NewPostgresWith(container.DSN, defaultConnPool())
		_ = source.AutoMigrate(&database.AnyEntity{})
		sut := database.NewGormTransactionManager(source)

		err := sut.Execute(context.Background(), func(tx domain.Transaction) error {
			if err := tx.GetTx().(*gorm.DB).Create(&database.AnyEntity{}).Error; err != nil {
				return err
			}
			return nil
		})
		var actual database.AnyEntity
		source.Find(&actual, 1)

		assert.NoError(t, err)
		assert.Equal(t, uint(1), actual.ID)

		t.Cleanup(func() {
			source.Unscoped().Delete(&actual)
		})
	})

	t.Run("Transaction Manager Rollback Test", func(t *testing.T) {
		source := NewPostgresWith(container.DSN, defaultConnPool())
		_ = source.AutoMigrate(&database.AnyEntity{})
		sut := database.NewGormTransactionManager(source)

		rollbackErr := sut.Execute(context.Background(), func(tx domain.Transaction) error {
			transaction := tx.GetTx().(*gorm.DB)
			transaction.Create(&database.AnyEntity{})
			transaction.Create(&database.AnyEntity{})
			return errors.New("rollback!!")
		})
		commitErr := sut.Execute(context.Background(), func(tx domain.Transaction) error {
			tx.GetTx().(*gorm.DB).Create(&database.AnyEntity{
				BaseModel: database.BaseModel{
					UpdatedAt: time.Now(),
				},
			})
			return nil
		})
		var actual []database.AnyEntity
		source.Find(&actual)
		marshal, _ := json.Marshal(actual)
		t.Log(string(marshal))

		assert.Error(t, rollbackErr)
		assert.NoError(t, commitErr)
		assert.Len(t, actual, 1)

		t.Cleanup(func() {
			source.Unscoped().Delete(&actual)
		})
	})
}
