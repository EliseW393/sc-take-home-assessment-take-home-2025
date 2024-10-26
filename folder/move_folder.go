package folder

import (
	"strings"
)

// TODO: errors, error messages 
// basic functionality: every path that has string, add dst before it (todo: other cases)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

	// iterate through every folder path
	for i := range f.Folders { // have to use i not _, otherwise modifying a copy
		// split path into array, check where name is, add dst before, reconstruct path
		pathLevels := strings.Split(f.Folders[i].Path, ".") // slice of path segments
		for j, element := range pathLevels {
			if element == name {
				pathLevels[j] = dst + "." + element // making it dst.name in the slice
				f.Folders[i].Path = strings.Join(pathLevels, ".") // put back together
				// stop here?
			}

	return []Folder{}, nil
}

// TODO: helper checking functions
