package objects

type Space struct {
	id        string
	owner     string
	resources []*Resource
}

func (space *Space) Id() string {
	return space.id
}

func (space *Space) Owner() string {
	return space.owner
}

func (space *Space) Resources() []*Resource {
	return space.resources
}

func (space *Space) AddResource(res *Resource) {
	space.resources = append(space.resources, res)
}

func NewSpace(id, owner string) *Space {
	return &Space{id, owner, []*Resource{}}
}
