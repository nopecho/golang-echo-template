package datasource

import (
	"github.com/nopecho/golang-template/test/testcontainer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgres(t *testing.T) {
	container := testcontainer.NewPostgresContainer()

	t.Run("TestNewPostgres", func(t *testing.T) {
		sut := NewPostgres(container.DSN, DefaultConnPool())

		var actual []string
		sut.Raw("SELECT 1").Scan(&actual)

		assert.Len(t, actual, 1)
	})
}
