package folder

import (
	"strings"
)

// TODO: errors, error messages 
// basic functionality: every path that has string, add dst before it (todo: other cases)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

	updatedFolders := make([]Folder, len(f.Folder)) // assign - non declared, local
	_ = copy(updatedFolders, f.Folders) // ignore return value of copy 

	// iterate through every folder path
	for i := range updatedFolders { // have to use i not _, otherwise modifying a copy
		copiedFolder := folder // making copy, not altering original file
		// split path into array, check where name is, add dst before, reconstruct path
		pathLevels := strings.Split(updatedFolders[i].Path, ".") // slice of path segments
		for j, element := range pathLevels {
			if element == name {
				pathLevels[j] = dst + "." + element // making it dst.name in the slice
				updatedFolders[i].Path = strings.Join(pathLevels, ".") // put back together
				// stop here?
			}

	return updatedFolders, nil
}

// TODO: helper checking functions
