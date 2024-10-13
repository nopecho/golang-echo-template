package domain

type AnyModel struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func NewAnyModel(name string) *AnyModel {
	return &AnyModel{
		Name: name,
	}
}
