package chrome

type DirectoryEntryCallback = func(DirectoryEntry)

type DirectoryEntry interface {
	IsFile() bool
	IsDirectory() bool
	Name() string
	FullPath() string
	FileSystem() interface{}
}
