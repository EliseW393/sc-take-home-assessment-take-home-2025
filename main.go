package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {

	//res := folder.GenerateData() 
	//folder.PrettyPrint(res)     
	//folder.WriteSampleData(res)  

	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	folderData := folder.GetSampleData() 
	folderDriver := folder.NewDriver(folderData)
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	fmt.Printf("\nFolders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)

	// Example of retrieving all child folders
	fmt.Printf("\nRetrieving all child folders for 'stunning-horridus' under orgID %s:\n", orgID)
	childFolders := folderDriver.GetAllChildFolders(orgID, "stunning-horridus")
	folder.PrettyPrint(childFolders)

	// test moving folders
	name := "sharing-rictor"
	dst := "ultimate-random" // testing informally 
	fmt.Printf("\nMoving folder '%s' to new destination '%s'...\n", name, dst)
	updatedFolders, err := folderDriver.MoveFolder(name, dst)
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	folder.PrettyPrint(updatedFolders)
}