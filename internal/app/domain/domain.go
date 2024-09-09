package domain

import "math/rand/v2"

type Domain struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewDomain(name string) *Domain {
	return &Domain{
		ID:   rand.Int64(),
		Name: name,
	}
}
