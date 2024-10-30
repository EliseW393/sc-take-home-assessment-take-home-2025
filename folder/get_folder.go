package folder

import "github.com/gofrs/uuid"
import (
	"fmt"
	"strings"
	"github.com/gofrs/uuid"
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

	res := f.GetFoldersByOrgID(orgID)
	if len(res) == 0 {
		// fmt.Println("Error: No folders found for the specified organisation ID")
		// above not part of specified functionality (?)
		return []Folder{} 
	}

	// 2: Create an empty slice to hold the final folders
	var finalFolders []Folder
	parentExists := false

	// 3: Iterate through each folder in res (filtered by orgID)
	for _, folder := range res {
		// Check if the folder's path contains the parent name 
		if strings.HasPrefix(folder.Paths, name+".") { 
			// append that folder to the slice, as it is a child 
			finalFolders = append(finalFolders, folder)
		}
	}
	// 7: Return slice of folders, children of given folder name
	return finalFolders
}


// Useful string methods found:
// strings.Contains: if string contains another substring
// strings.Index: position of substring within string
// strings.Split: split into slice of substrings by delimeter
