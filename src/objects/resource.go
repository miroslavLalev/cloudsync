package objects

type Resource struct {
	path  string
	space Space
}

func (resource *Resource) Path() string {
	return resource.path
}

func (resource *Resource) Space() Space {
	return resource.space
}
