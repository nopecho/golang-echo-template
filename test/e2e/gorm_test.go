package e2e

import (
	"github.com/goccy/go-json"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/internal/util/gorm/datasource"
	"github.com/nopecho/golang-template/test/testcontainer"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"slices"
	"sync"
	"testing"
)

var (
	container1 = testcontainer.NewPostgresContainer()
	container2 = testcontainer.NewPostgresContainer()
	db1        = datasource.NewPostgresWith(container1.DSN, nil)
	db2        = datasource.NewPostgresWith(container2.DSN, nil)
)

func TestGorm(t *testing.T) {
	datasource.AutoMigrate(db1, &database.AnyEntity{})
	datasource.AutoMigrate(db2, &database.Any2Entity{})

	t.Run("save test", func(t *testing.T) {
		err := db2.Transaction(func(tx *gorm.DB) error {
			tx.Create(&database.Any2Entity{
				Payload: map[string]interface{}{
					"save": "test",
				},
			})
			return nil
		})
		if err != nil {
			t.Log(err)
		}

		var actual []database.Any2Entity
		db2.Find(&actual)
		str, _ := json.Marshal(actual)
		t.Log(string(str))

		assert.Len(t, actual, 1)

		t.Cleanup(func() {
			db2.Unscoped().Delete(&actual)
		})
	})

	t.Run("multi datasource test", func(t *testing.T) {
		var actual1 []string
		db1.Raw(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`).Scan(&actual1)
		assert.True(t, slices.Contains(actual1, "domain"))

		var actual2 []string
		db2.Raw(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`).Scan(&actual2)
		assert.True(t, slices.Contains(actual2, "domain2"))
	})

	t.Run("go insert test", func(t *testing.T) {
		size := 100
		wg := &sync.WaitGroup{}
		wg.Add(size)
		for i := range size {
			go func() {
				defer wg.Done()
				db2.Create(&database.Any2Entity{
					Payload: database.Jsonb{
						"batch": i,
					},
				})
			}()
		}
		wg.Wait()

		var count int64
		db2.Model(&database.Any2Entity{}).Count(&count)

		assert.Equal(t, int64(size), count)

		t.Cleanup(func() {
			db2.Where("1=1").Delete(&database.Any2Entity{})

			var deleted int64
			db2.Model(&database.Any2Entity{}).Count(&deleted)

			assert.Equal(t, int64(0), deleted)
		})
	})

	t.Run("batch insert test", func(t *testing.T) {
		size := 100
		entities := make([]database.AnyEntity, size)
		for i := range size {
			entities[i] = database.AnyEntity{}
		}

		db1.CreateInBatches(entities, 100)

		var count int64
		db1.Model(&database.AnyEntity{}).Count(&count)

		assert.Equal(t, int64(size), count)

		t.Cleanup(func() {
			db1.Where("1=1").Delete(&database.AnyEntity{})

			var deleted int64
			db1.Model(&database.AnyEntity{}).Count(&deleted)

			assert.Equal(t, int64(0), deleted)
		})
	})

	t.Cleanup(func() {
		container1.Terminate()
		container2.Terminate()
	})
}
