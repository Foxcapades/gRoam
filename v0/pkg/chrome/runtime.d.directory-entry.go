package chrome

type DirectoryEntryCallback = func(DirectoryEntry)

type DirectoryEntry interface {
	// A Boolean which is true if the entry represents a file. If it's not a file,
	// this value is false.
	IsFile() bool

	// A Boolean which is true if the entry represents a directory; otherwise,
	// it's false.
	IsDirectory() bool

	// A USVString containing the name of the entry (the final part of the path,
	// after the last "/" character).
	Name() string

	// A USVString object which provides the full, absolute path from the file
	// system's root to the entry; it can also be thought of as a path which is
	// relative to the root directory, prepended with a "/" character.
	FullPath() string

	// A FileSystem object representing the file system in which the entry is
	// located.
	Filesystem() interface{}
}
