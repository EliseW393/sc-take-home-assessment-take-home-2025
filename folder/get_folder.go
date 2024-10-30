package folder

import (
	"strings"
	"github.com/gofrs/uuid"
	"regexp"
    "log"
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
func (f *driver) GetAllChildFolders(orgID uuid.UUID, parentFolderName string) []Folder { 

	var pathRegex = regexp.MustCompile(`^([a-zA-Z0-9-_]+)(\.[a-zA-Z0-9-_]+)*$`)
	res := f.GetFoldersByOrgID(orgID)

	if orgID == uuid.Nil {
        log.Println("Error: invalid orgID (nil)")
        return []Folder{}
    } else if len(res) == 0 {
        log.Println("Warning: No folders found for the specified orgID")
        return []Folder{}
    }

	// Slice to store child folders 
	var finalFolders []Folder

	// Iterate over each folder in filtered results (res)
	for _, folder := range res {

		if !pathRegex.MatchString(folder.Paths) {
            log.Printf("Warning: Invalid path structure in folder %s with path %s\n", folder.Name, folder.Paths)
            continue // skippng folders with invalid paths 
        }

		// Check if the folder's path contains the specified parent name
		if strings.Contains(folder.Paths, parentFolderName+".") { 
			// Append matching folder to finalFolders as it is a child of the specified folder 
			finalFolders = append(finalFolders, folder)
		}
	}

	return finalFolders
}