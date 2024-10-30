package folder

import (
	"errors"
	"strings"
	"regexp"
	"github.com/gofrs/uuid"
)

func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

	if name == "" || dst == "" {
		emptyNameError := errors.New("Error: Source and destination folder names must not be empty\n")
		return f.folders, emptyNameError
	}

	// functionality: if source and dst are the same
	if name == dst {
		sourceAsDestError := errors.New("Error: Cannot move a folder to itself\n")
		return f.folders, sourceAsDestError
	}

	// validPath := regexp.MustCompile(`^([a-zA-Z0-9-]+)(\.[a-zA-Z0-9-]+)*$`) // turns it into object to verify format

	updatedFolders := make([]Folder, len(f.folders)) // assign - non declared, local
	_ = copy(updatedFolders, f.folders) // ignore return value of copy 

	var sourceOrgID uuid.UUID
	var dstOrgID uuid.UUID 
	sourceFound := false
	destinationLocated := false

	for _, folder := range updatedFolders {
		if folder.Name == name {
			sourceOrgID = folder.OrgId
			sourceFound = true
		}

		if folder.Name == dst {//&& updatedFolders[i].OrgId == sourceOrgID {
			dstOrgID = folder.OrgId 
			destinationLocated = true
		}
	}

	if !sourceFound { 
		sourceNotExistError := errors.New("Error: Source folder does not exist\n")
		return f.folders, sourceNotExistError
	}

	if !destinationLocated { // Not working - need to actually check name of folder too 
		dstNotExistError := errors.New("Error: Destination folder does not exist\n")
		return f.folders, dstNotExistError
	}

	// if dst org ID doesnt match that of source - check orgID of folder where name == dst
	if dstOrgID != sourceOrgID {
		diffOrgError := errors.New("Error: Cannot move a folder to a different organization\n")
		return f.folders, diffOrgError
	}

	// iterate through every folder path
	for i := range updatedFolders { // have to use i not _, otherwise modifying a copy
		
		// if !validPath.MatchString(updatedFolders[i].Paths) {
			// invalidPathStructureError := errors.New("Error: Invalid folder path structure\n")
			// return f.folders, invalidPathStructureError
		// }

		// split path into array, check where name is, add dst before, reconstruct path
		pathLevels := strings.Split(updatedFolders[i].Paths, ".") // slice of path segments - check source exists first? save looping

		for j, folder := range pathLevels { // shouldnt iterate if its not there to begin with - check 
			if folder == name { // pathLevels at j is equal to the name (source)

				children := pathLevels[j+1:] // Only segments after `name`
            	if contains(children, dst) {
					childOfItselfError := errors.New("Error: Cannot move a folder to a child of itself\n")
                	return f.folders, childOfItselfError
            	}

				pathLevels[j] = dst + "." + folder // making it dst.name in the slice
				updatedFolders[i].Paths = strings.Join(pathLevels, ".") // put back together
				break
			}
		}
	}

	return updatedFolders, nil
}	

