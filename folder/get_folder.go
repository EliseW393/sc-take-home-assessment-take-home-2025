package folder

import "github.com/gofrs/uuid"
import (
	// "fmt"
	"strings"
	// "github.com/gofrs/uuid"
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

// GetAllChildFolders retrieves all child folders of a specified folder name within an organisation.
// It filters folders by the given organization ID and returns those whose paths begin with the specified name.
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder { 

	res := f.GetFoldersByOrgID(orgID)
	if len(res) == 0 {
		// If no folders are found for the orgID, return empty slice 
		return []Folder{} 
	}

	// Slice to store child folders 
	var finalFolders []Folder

	// Iterate over each folder in filtered results (res)
	for _, folder := range res {
		// Check if the folder's path starts with the specified parent name
		if strings.HasPrefix(folder.Paths, name+".") { 
			// Append matching folder to finalFolders as it is a child of the specified folder 
			finalFolders = append(finalFolders, folder)
		}
	}

	return finalFolders
}