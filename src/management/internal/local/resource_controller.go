package local

import (
	"bufio"

	o "github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/repository"
	"github.com/miroslavLalev/cloudsync/src/util"
)

func NewResourceController() *ResourceController {
	return &ResourceController{facade: repository.GetRepositoryFacade()}
}

type ResourceController struct {
	facade repository.RepositoryFacade
}

func (rc *ResourceController) Get(id string) (*o.Resource, error) {
	fi, err := rc.facade.GetFile(util.DecodePath(id))
	if err != nil {
		return nil, err
	}

	res := o.NewResource(id, fi.Path(), deduceType(fi))
	return res, nil
}

func (rc *ResourceController) Create(resource o.Resource) {

}

func (rc *ResourceController) Upload(resource o.Resource, content bufio.Reader) {

}

func (rc *ResourceController) Delete(resource o.Resource) {

}

func (rc *ResourceController) Share(resource o.Resource, other o.Space) {

}
