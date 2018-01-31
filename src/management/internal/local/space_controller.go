package local

import (
	"github.com/miroslavLalev/cloudsync/src/db"
	o "github.com/miroslavLalev/cloudsync/src/objects"
	"github.com/miroslavLalev/cloudsync/src/repository"
	"github.com/miroslavLalev/cloudsync/src/repository/file"
	"github.com/miroslavLalev/cloudsync/src/util"
)

func NewSpaceController() *SpaceController {
	return &SpaceController{facade: repository.GetRepositoryFacade()}
}

type SpaceController struct {
	facade repository.RepositoryFacade
}

func (uc *SpaceController) Create(user string) error {
	space := o.NewSpace(user, user)
	err := db.GetController().AddSpace(space)
	if err != nil {
		return err
	}

	err = uc.facade.CreateDir(space.Id())
	if err != nil {
		return err
	}

	return nil
}

func (uc *SpaceController) Delete(space *o.Space) error {
	err := db.GetController().DeleteSpace(space.Owner())
	if err != nil {
		return err
	}

	return uc.facade.Remove(space.Id())
}

func (uc *SpaceController) Get(id string) (*o.Space, error) {
	files, err := uc.facade.ListContent(id)
	if err != nil {
		return nil, err
	}

	space := o.NewSpace(id, "")
	for _, fi := range files {
		space.AddResource(o.NewResource(util.EncodePath(fi.Path()), fi.Path(), deduceType(fi)))
	}

	return space, nil
}

func deduceType(fi *file.File) o.Type {
	if fi.IsDir() {
		return o.Dir
	}

	return o.File
}
