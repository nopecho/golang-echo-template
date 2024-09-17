package domain

type AnyCreateCommand struct {
	Name string
}

type AnyUpdateCommand struct {
	ID   uint64
	Name string
}

type AnyQuery struct {
	ID int64
}
