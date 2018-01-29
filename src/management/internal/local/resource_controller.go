package local

import (
	"bufio"

	o "github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/repository"
)

func NewResourceController() *ResourceController {
	return &ResourceController{facade: repository.GetRepositoryFacade()}
}

type ResourceController struct {
	facade repository.RepositoryFacade
}

func (rc *ResourceController) Create(resource o.Resource) {

}

func (rc *ResourceController) Upload(resource o.Resource, content bufio.Reader) {

}

func (rc *ResourceController) Delete(resource o.Resource) {

}

func (rc *ResourceController) Share(resource o.Resource, other o.Space) {

}
