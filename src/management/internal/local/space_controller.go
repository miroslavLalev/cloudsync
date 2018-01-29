package local

import (
	o "github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/repository"
)

func NewSpaceController() *SpaceController {
	return &SpaceController{facade: repository.GetRepositoryFacade()}
}

type SpaceController struct {
	facade repository.RepositoryFacade
}

func (uc *SpaceController) Create(space o.Space) {
	uc.facade.CreateDir(space.Id())
}

func (uc *SpaceController) Delete(space o.Space) {
	uc.facade.Remove(space.Id())
}
