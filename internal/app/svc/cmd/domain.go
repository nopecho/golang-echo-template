package cmd

type DomainCreateCommand struct {
	Name string
}

type DomainUpdateCommand struct {
	ID   int64
	Name string
}
