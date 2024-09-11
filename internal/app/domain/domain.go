package domain

type Domain struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewDomain(name string) *Domain {
	return &Domain{
		Name: name,
	}
}
