package folder

import (
	"strings"
)

// consider - cant be under its own child!

// TODO: errors, error messages 
// basic functionality: every path that has string, add dst before it (todo: other cases)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

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
