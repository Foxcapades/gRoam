package x

type DirectoryEntry struct {
	isFile     bool
	name       string
	fullPath   string
	fileSystem interface{}
}

func (d *DirectoryEntry) IsFile() bool {
	return d.isFile
}

func (d *DirectoryEntry) IsDirectory() bool {
	return !d.isFile
}

func (d *DirectoryEntry) Name() string {
	return d.name
}

func (d *DirectoryEntry) FullPath() string {
	return d.fullPath
}

func (d *DirectoryEntry) Filesystem() interface{} {
	return d.fileSystem
}
