package svc

type DomainCreateCommand struct {
	Name string
}

type DomainUpdateCommand struct {
	ID   uint64
	Name string
}

type DomainQuery struct {
	ID int64
}
