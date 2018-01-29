package repository

import (
	"fmt"
	"os"

	"github.com/miroslavLalev/cloudsync/src/repository/internal/local"
)

const repoPathVariable string = "REPO"

// TODO: consider generic naming like "sections" and "resources" (abstract away from storage)
type RepositoryFacade interface {
	CreateDir(uri string)

	CreateFile(uri string)

	Remove(uri string)

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
