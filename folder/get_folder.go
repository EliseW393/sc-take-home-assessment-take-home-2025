package folder

import "github.com/gofrs/uuid"
import (
	"strings" // string manipulation functions
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder { 
	// (Note; Don't forget function description!)
	// TODO: Implement logic to get all child folders of given parent, with orgID

	// 1: Retrieve all folders for the given orgID
		// - Call GetFoldersByOrgID on f and pass in the orgID argument.
		// - Store the result in a variable, res.
	res := f.GetFoldersByOrgID(orgID)

	// 2: Create an empty slice to hold the final folders
	var finalFolders []Folder

	// 3: Iterate through each folder in res (filtered by orgID)
	for _, folder := range res {
		// Check if the folder's path contains the parent name (string.Contains)
		if strings.HasPrefix(folder.Paths, name+".") { 
			// append that folder to the slice, as it is a child 
			finalFolders = append(finalFolders, folder)
		}
	}
		}
	// 7: Return slice of folders, children of given folder name
	return finalFolders
}

// Useful string methods found:
// strings.Contains: if string contains another substring
// strings.Index: position of substring within string
// strings.Split: split into slice of substrings by delimeter
