package management

import o "github.com/miroslavLalev/cloudsync/src/objects"

type SpaceController interface {
	Create(space o.Space)
	Delete(space o.Space)
}
