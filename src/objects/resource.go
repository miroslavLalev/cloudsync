package objects

type Type string

const (
	File Type = "FILE"
	Dir  Type = "DIR"
)

type Resource struct {
	id        string
	name      string
	rtype     Type
	resources []*Resource
}

func (resource *Resource) Id() string {
	return resource.id
}

func (resource *Resource) Name() string {
	return resource.name
}

func (resource *Resource) Type() Type {
	return resource.rtype
}

func (resource *Resource) Resources() []*Resource {
	return resource.resources
}

func (resource *Resource) AddResource(r *Resource) {
	resource.resources = append(resource.resources, r)
}

func NewResource(id, name string, rtype Type) *Resource {
	return &Resource{
		id:        id,
		name:      name,
		rtype:     rtype,
		resources: []*Resource{},
	}
}
