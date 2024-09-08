package examples

import (
	"github.com/google/uuid"
	"github.com/nopecho/golang-template/pkg/synk"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Benchmark(b *testing.B) {
	size := 100000
	b.Run("uuid v7", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			cm := synk.OpenBufferChannel[string](size)

			for range size {
				go func() {
					v7, _ := uuid.NewV7()
					cm.Send(v7.String())
				}()
			}
			result := cm.WaitReceive()

			assert.Len(b, result, size)
		}
	})

	b.Run("uuid v4", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			cm := synk.OpenBufferChannel[string](size)

			for range size {
				go func() {
					id := uuid.NewString()
					cm.Send(id)
				}()
			}
			result := cm.WaitReceive()

			assert.Len(b, result, size)
		}
	})
}
