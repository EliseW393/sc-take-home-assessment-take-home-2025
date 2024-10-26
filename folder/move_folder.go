package folder

import (
	"errors"
	"strings"
)

//moveFolder("bravo", "charlie")
// Error: Cannot move a folder to a child of itself

//moveFolder("bravo", "bravo")
// Error: Cannot move a folder to itself

//moveFolder("bravo", "foxtrot")
// Error: Cannot move a folder to a different organization

//moveFolder("invalid_folder", "delta")
// Error: Source folder does not exist

//moveFolder("bravo", "invalid_folder")
// Error: Destination folder does not exist

// TODO: errors, error messages 
// basic functionality: every path that has string, add dst before it (todo: other cases)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

	// to implement
	childOfItselfError := errors.New("Error: Cannot move a folder to a child of itself")
	// functionality: check if dst is a child of that current folder 
	// i.e., if the names AFTER the source INCLUDE the dst, then this occurs 

	folderToItselfError := errors.New("Error: Cannot move a folder to itself")
	// functionality: if source and dst are the same

	diffOrgError := errors.New("Error: Cannot move a folder to a different organization")
	// if dst org ID doesnt match that of source - check orgID of folder where name == dst

	sourceNotExistError := errors.New("Error: Source folder does not exist")
	// if checked all paths and no changes made, source not in any path, DNE

	dstNotExistError := errors.New("Error: Destination folder does not exist")
	// check all folder names to make sure dest folder exists before inserting anywhere 

	// basic functionality 
	updatedFolders := make([]Folder, len(f.folders)) // assign - non declared, local
	_ = copy(updatedFolders, f.folders) // ignore return value of copy 

	// iterate through every folder path
	for i := range updatedFolders { // have to use i not _, otherwise modifying a copy
		// split path into array, check where name is, add dst before, reconstruct path
		pathLevels := strings.Split(updatedFolders[i].Paths, ".") // slice of path segments
		for j, element := range pathLevels {
			if element == name {
				pathLevels[j] = dst + "." + element // making it dst.name in the slice
				updatedFolders[i].Paths = strings.Join(pathLevels, ".") // put back together
				// stop here?
			}
		}
	}

	return updatedFolders, nil
}

// TODO: helper checking functions
