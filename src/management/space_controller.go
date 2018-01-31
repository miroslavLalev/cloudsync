package management

import (
	o "github.com/miroslavLalev/cloudsync/src/objects"
)

type SpaceController interface {
	Create(user string) error
	Delete(space *o.Space) error
	Get(id string) (*o.Space, error)
}
