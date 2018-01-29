package objects

type Space struct {
	id    string
	owner string
	// TODO: multiple owners
	// sharedResources map[string][]string
}

func (space *Space) Id() string {
	return space.id
}

func (space *Space) Owner() string {
	return space.owner
}

func NewSpace(id, owner string) *Space {
	return &Space{id, owner}
}
