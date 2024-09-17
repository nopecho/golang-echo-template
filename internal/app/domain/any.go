package domain

type AnyModel struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func NewAnyModel(name string) *AnyModel {
	return &AnyModel{
		Name: name,
	}
}
