package local

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/miroslavLalev/cloudsync/src/repository/file"
)

type Repository struct {
	rootPath string
}

func (r Repository) CreateDir(uri string) error {
	return os.Mkdir(r.buildPath(uri), 0664)
}

func (r Repository) Remove(uri string) error {
	return os.Remove(r.buildPath(uri))
}

func (r Repository) CreateFile(uri string) {
	os.Create(r.buildPath(uri))
}

func (r Repository) GetFile(uri string) (*file.File, error) {
	f, err := os.Stat(r.buildPath(uri))
	return warpFile(f), err
}

func warpFile(f os.FileInfo) *file.File {
	return file.NewFile(f.Name(), filepath.Base(f.Name()), f.IsDir())
}

func (r Repository) ListContent(uri string) ([]*file.File, error) {
	files, err := ioutil.ReadDir(r.buildPath(uri))
	return wrapFiles(files), err
}

func wrapFiles(files []os.FileInfo) []*file.File {
	result := []*file.File{}
	for _, f := range files {
		result = append(result, warpFile(f))
	}

	return result
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
	return r.rootPath + "/" + uri
}

func NewRepository(rootPath string) *Repository {
	return &Repository{rootPath}
}
