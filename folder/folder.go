package folder

import "github.com/gofrs/uuid"

type IDriver interface { // any type that implements IDriver must have these methods - what functions should exist
	// GetFoldersByOrgID returns all folders that belong to a specific orgID.
	GetFoldersByOrgID(orgID uuid.UUID) []Folder
	// component 1
	// Implement the following methods:
	// GetAllChildFolders returns all child folders of a specific folder.
	GetAllChildFolders(orgID uuid.UUID, name string) []Folder

	// component 2
	// Implement the following methods:
	// MoveFolder moves a folder to a new destination.
	MoveFolder(name string, dst string) ([]Folder, error)
}

type driver struct { // implements IDriver interface, storing at slide of Folder structs. Where folder data lives
	// define attributes here
	// data structure to store folders
	// or preprocessed data

	// example: feel free to change the data structure, if slice is not what you want
	folders []Folder // slice: dynamic array, here holds Folder structs
}

func NewDriver(folders []Folder) IDriver {
	return &driver{
		// initialize attributes here
		folders: folders, // store folder data inside driver instance
	}
}
