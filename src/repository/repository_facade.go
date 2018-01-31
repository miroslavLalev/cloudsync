package repository

import (
	"fmt"
	"os"

	"github.com/miroslavLalev/cloudsync/src/repository/file"
	"github.com/miroslavLalev/cloudsync/src/repository/internal/local"
)

const repoPathVariable string = "REPO"

type RepositoryFacade interface {
	CreateDir(uri string) error

	CreateFile(uri string)

	ListContent(uri string) ([]*file.File, error)

	Remove(uri string) error

	UploadFile(uri string, content []byte)

	DownloadFile(uri string)

	CreateLink(from, to string)
}

func GetRepositoryFacade() RepositoryFacade {
	return facade
}

var facade RepositoryFacade

func init() {
	repoPath := os.Getenv(repoPathVariable)
	if repoPath == "" {
		panic(fmt.Sprintf("Failed to read repository path from variable '%v'", repoPathVariable))
	}

	//TODO: remote repo
	facade = RepositoryFacade(local.NewRepository(repoPath))
}
