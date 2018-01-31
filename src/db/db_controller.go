package db

import (
	"github.com/miroslavLalev/cloudsync/src/db/internal"
	"github.com/miroslavLalev/cloudsync/src/objects"
)

type Controller interface {
	AddSpace(space *objects.Space) error
	GetSpace(user string) (string, error)
	DeleteSpace(user string) error
}

func GetController() Controller {
	return Controller(internal.GetController())
}
