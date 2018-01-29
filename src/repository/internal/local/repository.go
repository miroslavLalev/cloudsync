package local

import (
	"fmt"
	"os"
)

type Repository struct {
	rootPath string
}

func (r Repository) CreateDir(uri string) {
	os.Mkdir(r.buildPath(uri), 0440)
}

func (r Repository) Remove(uri string) {
	os.Remove(r.buildPath(uri))
}

func (r Repository) CreateFile(uri string) {
	// TODO: 0644 instead 0666
	os.Create(r.buildPath(uri))
}

func (r Repository) UploadFile(uri string, content []byte) {
	// TODO: use Repository#Create
	file, err := os.Create(r.buildPath(uri))
	if err != nil {
		// TODO: err
		fmt.Println("Error while creating file")
		return
	}

	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Error writing to file")
	}

	file.Close()
}

func (r Repository) DownloadFile(uri string) {
	// TODO: implement (return []byte)
}

func (r Repository) CreateLink(from, to string) {
	os.Symlink(r.buildPath(from), r.buildPath(to))
}

func (r Repository) buildPath(uri string) string {
	return r.rootPath + uri
}

func NewRepository(rootPath string) *Repository {
	return &Repository{rootPath}
}
