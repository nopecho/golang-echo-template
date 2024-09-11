package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRedisContainer(t *testing.T) {
	sut := NewRedisContainer()

	t.Run("NewRedisContainer", func(t *testing.T) {
		actual := sut.Endpoint

		assert.Contains(t, actual, "localhost:")
	})

	t.Cleanup(func() {
		sut.Terminate()
	})
}
