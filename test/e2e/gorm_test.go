package e2e

import (
	"encoding/json"
	"github.com/nopecho/golang-template/internal/app/infra/database"
	"github.com/nopecho/golang-template/pkg/gorm/datasource"
	"github.com/nopecho/golang-template/test"
	"github.com/stretchr/testify/assert"
	"github.com/thoas/go-funk"
	"gorm.io/gorm"
	"testing"
)

var (
	container1 = test.NewPostgresContainer()
	container2 = test.NewPostgresContainer()
)

func TestGorm(t *testing.T) {
	db1 := datasource.NewPostgres(container1.DSN, datasource.DefaultConnPool())
	db2 := datasource.NewPostgres(container2.DSN, datasource.DefaultConnPool())
	datasource.AutoMigrate(db1, &database.DomainEntity{})
	datasource.AutoMigrate(db2, &database.Domain2Entity{})

	t.Run("save test", func(t *testing.T) {
		err := db2.Transaction(func(tx *gorm.DB) error {
			tx.Create(&database.Domain2Entity{
				Payload: map[string]interface{}{
					"save": "test",
				},
			})
			return nil
		})
		if err != nil {
			t.Log(err)
		}

		var actual []database.Domain2Entity
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
		assert.True(t, funk.Some(actual1, func(s string) bool {
			return s == "domain_entities"
		}))

		var actual2 []string
		db2.Raw(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';`).Scan(&actual2)
		assert.True(t, funk.Some(actual2, func(s string) bool {
			return s == "domain2_entities"
		}))
	})

	t.Cleanup(func() {
		container1.Terminate()
		container2.Terminate()
	})
}
