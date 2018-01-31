package file

type File struct {
	path  string
	name  string
	isDir bool
}

func (f *File) Path() string {
	return f.path
}

func (f *File) Name() string {
	return f.name
}

func (f *File) IsDir() bool {
	return f.isDir
}

func NewFile(path, name string, isDir bool) *File {
	return &File{path: path, name: name, isDir: isDir}
}
