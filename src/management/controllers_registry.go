package management

import (
	"github.com/miroslavLalev/cloudsync/src/management/internal/local"
)

type ControllersRegistry struct {
	sc SpaceController
	rc ResourceController
}

func (r *ControllersRegistry) GetSpaceController() SpaceController {
	return r.sc
}

func GetRegistry() *ControllersRegistry {
	return registry
}

var registry *ControllersRegistry

func init() {
	//TODO: add remote
	registry = &ControllersRegistry{
		sc: SpaceController(local.NewSpaceController()),
		rc: ResourceController(local.NewResourceController()),
	}
}
