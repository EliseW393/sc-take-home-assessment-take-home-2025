package folder

import "github.com/gofrs/uuid"
import (
	"strings" // string manipulation functions
	"encoding/json" // working with JSON
)

func GetAllFolders() []Folder { 
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder { 
	folders := f.folders 

	res := []Folder{}
	for _, f := range folders { 
		if f.OrgId == orgID { 
			res = append(res, f) 
		}
	}

	return res

}

func (f *driver) GetUniqueChildren(path string, uniqueChildren *[]string) {
	// (Note; Don't forget function description!)
	// TODO: Implement logic to extract unique children from string
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder { 
	// (Note; Don't forget function description!)
	// TODO: Implement logic to get all child folders of given parent, with orgID

	// 1: Retrieve all folders for the given orgID
		// - Call GetFoldersByOrgID on f and pass in the orgID argument.
		// - Store the result in a variable, res.

	// 2: Create an empty slice to hold the final folders

	// 3: Create an empty slice to hold unique child names

	// 4: Iterate through each folder in res (filtered by orgID)
		// Check if the folder's path contains the parent name (string.Contains)
			// Truncate the right string (may be empty)

			// 5: Call the helper function to find unique children
        	// Use pointer to the unique child name slice
        	// If no children are found, return 
        	// If children are found, split and append them to slice if unique

	// 6: Iterate through the unique child name list
		// Append matching folders from res to finalFolders based on the child name

	// 7: Return slice of folders, children of given folder name
	return []Folder{}
}

// Useful string methods found:
// strings.Contains: if string contains another substring
// strings.Index: position of substring within string
// strings.Split: split into slice of substrings by delimeter
