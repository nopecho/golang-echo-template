package http

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttpClient(t *testing.T) {
	mock := "https://jsonplaceholder.typicode.com"
	sut := NewClient(mock, nil)

	t.Run("GetJson", func(t *testing.T) {
		res, _ := sut.GetJson("/posts/1")

		actual, _ := json.Marshal(res)
		t.Log(string(actual))

		assert.Equal(t, 1.0, res["id"])
	})

	t.Run("GetRaw", func(t *testing.T) {
		res, _ := sut.GetRaw("/posts/1")

		t.Log(string(res))

		assert.NotNil(t, res)
	})

	t.Run("PostJson", func(t *testing.T) {
		data := ANY{
			"title":  "foo",
			"body":   "bar",
			"userId": 1,
		}
		res, _ := sut.PostJson("/posts", data)

		actual, _ := json.Marshal(res)
		t.Log(string(actual))

		assert.Equal(t, "foo", res["title"])
	})

	t.Run("PatchJson", func(t *testing.T) {
		data := ANY{
			"title": "bar",
		}
		res, _ := sut.PatchJson("/posts/1", data)

		actual, _ := json.Marshal(res)
		t.Log(string(actual))

		assert.Equal(t, "bar", res["title"])
	})
}
