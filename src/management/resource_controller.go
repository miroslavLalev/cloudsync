package management

import (
	"bufio"

	o "github.com/miroslavLalev/cloudsync/src/objects"
)

type ResourceController interface {
	Create(resource o.Resource)
	Upload(resource o.Resource, content bufio.Reader)
	Delete(resource o.Resource)
	Share(resource o.Resource, other o.Space)
}
