package cmd

type DomainCreateCommand struct {
	Name string
}

type DomainUpdateCommand struct {
	ID   uint64
	Name string
}
