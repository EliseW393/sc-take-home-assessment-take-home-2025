package folder

import (
	"errors"
	"strings"
	"github.com/gofrs/uuid"
	"regexp"
    "log"
	"fmt"
)

// contains is used to check for an item in a slice of strings 
// (i.e., if a destination folder exists in the list of children)
func contains(slice []string, item string) bool {
    for _, v := range slice {
        if v == item {
            return true
        }
    }
    return false
}

// MoveFolder moves a specified source folder to a new destination folder within the same organisation.
// It performs validation checks on the source and destination folders to make sure they exist, are in the same
// organisation, and are not in a parent-child loop/cycle.
func (f *driver) MoveFolder(source string, dst string) ([]Folder, error) {

	var pathRegex = regexp.MustCompile(`^([a-zA-Z0-9-_]+)(\.[a-zA-Z0-9-_]+)*$`)

	if source == "" || dst == "" {
		emptyNameError := errors.New("Error: Source and destination folder names must not be empty\n")
		return f.folders, emptyNameError
	} else if source == dst {
		sourceAsDestError := errors.New("Error: Cannot move a folder to itself\n")
		return f.folders, sourceAsDestError
	}

	var sourceOrgID, dstOrgID uuid.UUID
	sourceFound, destinationLocated := false, false

	// Locate source and destination folders to validate existence, obtain OrgIDs
	for _, folder := range f.folders {

		if !pathRegex.MatchString(folder.Paths) {
			log.Printf("Warning: Invalid path structure in folder %s with path %s\n", folder.Name, folder.Paths)
			continue // skipping any folder with invalid path format
		}

		if folder.Name == source {
			sourceOrgID = folder.OrgId
			sourceFound = true
		} else if folder.Name == dst {
			dstOrgID = folder.OrgId 
			destinationLocated = true
		}
	}

	if !sourceFound { 
		sourceNotExistError := errors.New("Error: Source folder does not exist\n")
		return f.folders, sourceNotExistError
	} else if !destinationLocated { 
		dstNotExistError := errors.New("Error: Destination folder does not exist\n")
		return f.folders, dstNotExistError
	} else if dstOrgID != sourceOrgID {
		diffOrgError := errors.New("Error: Cannot move a folder to a different organisation\n")
		return f.folders, diffOrgError
	}

	// Iterate through each folder path to update parent of source
	for i := range f.folders { 

		// Split folder path into parts (individual folders) to check for child moves and apply changes
		pathLevels := strings.Split(f.folders[i].Paths, ".") 
		fmt.Printf("Checking folder: %s with path: %s\n", f.folders[i].Name, f.folders[i].Paths)

		// Find location of source folder in path (if exists)
		for j, folder := range pathLevels { // Iterating through path folders
			if folder == source {

				// Checking if any parent in pathLevels up to `source` contains `dst`
				parents := pathLevels[:j] 
				if contains(parents, dst) {
					parentError := errors.New("Error: Destination folder is already a parent of the source folder, cannot move.\n")
					return f.folders, parentError
				}

				// Checking if folders after source folder includes dst
				children := pathLevels[j+1:] 
				fmt.Printf("CHILDREN OF %s: %v, NO %s\n", source, children, dst)
				fmt.Println("RESULT: ", contains(children, dst))
				if contains(children, dst) {
					//childOfItselfError := errors.New("Error: Cannot move a folder to a child of itself\n")
					return f.folders, errors.New("Error: Cannot move a folder to a child of itself\n")
				}
				
				// Update path of folder to reflect new destination (parent)
				pathLevels[j] = dst + "." + folder 
				f.folders[i].Paths = strings.Join(pathLevels, ".") 
				break
			}
		}
	}
	return f.folders, nil
}	

