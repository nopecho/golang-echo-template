package testcontainer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostgresContainer(t *testing.T) {
	sut := NewPostgresContainer()

	t.Run("NewPostgresContainer", func(t *testing.T) {
		actual := sut.DSN

		assert.Contains(t, actual, "host=localhost user=test password=test dbname=test")
	})

	t.Cleanup(func() {
		sut.Terminate()
	})
}
