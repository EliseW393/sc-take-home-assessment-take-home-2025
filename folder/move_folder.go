package folder

import (
	"errors"
	"strings"
	"fmt"
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

// TODO: errors, error messages 
// basic functionality: every path that has string, add dst before it (todo: other cases)
func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {
	// (TODO: function description)

	var sourceOrgID uuid.UUID
	var dstOrgID uuid.UUID

	// functionality: if source and dst are the same
	if name == dst {
		folderToItselfError := errors.New("Error: Cannot move a folder to itself\n")
		return f.folders, folderToItselfError
	}

	updatedFolders := make([]Folder, len(f.folders)) // assign - non declared, local
	_ = copy(updatedFolders, f.folders) // ignore return value of copy 

	sourceFound := false
	destinationLocated := false
	for i := range updatedFolders {
		if updatedFolders[i].Name == name {
			sourceOrgID = updatedFolders[i].OrgId
			sourceFound = true
		}

		if updatedFolders[i].Name == dst {//&& updatedFolders[i].OrgId == sourceOrgID {
			dstOrgID = updatedFolders[i].OrgId
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
		dstOrgID := updatedFolders[i].OrgId

		fmt.Printf("Type of dstOrgID: %T, Value: %v\n", dstOrgID, dstOrgID)  // checking 

		// split path into array, check where name is, add dst before, reconstruct path
		pathLevels := strings.Split(updatedFolders[i].Paths, ".") // slice of path segments - check source exists first? save looping
		for j, element := range pathLevels { // shouldnt iterate if its not there to begin with - check 
			if element == name { // pathLevels at j is equal to the name (source)

				// check all children below source - make sure dst isnt there

				children := pathLevels[0:len(pathLevels)-1]
				// functionality: check if dst is a child of that current folder
				if contains(children, dst) { // too many nextes loops, fix 
					childOfItselfError := errors.New("Error: Cannot move a folder to a child of itself\n")
					return f.folders, childOfItselfError
				}

				pathLevels[j] = dst + "." + element // making it dst.name in the slice
				updatedFolders[i].Paths = strings.Join(pathLevels, ".") // put back together
				break
			}
		}
	}



	return updatedFolders, nil
}

