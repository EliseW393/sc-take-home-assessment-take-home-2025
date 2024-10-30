package main

import (
	"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {

	// Generate new sample data
	res := folder.GenerateData() // Generate fresh folder data
	folder.PrettyPrint(res)      // Print the generated data
	folder.WriteSampleData(res)  // Write the data to sample.json

	// Use the DefaultOrgID to filter folders
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)

	// Get all folders (from the sample data)
	folderData := folder.GetSampleData() // Load sample data from sample.json

	// Create a new folder driver
	folderDriver := folder.NewDriver(folderData)

	// Get folders by organisation ID
	orgFolder := folderDriver.GetFoldersByOrgID(orgID)

	fmt.Printf("\nFolders for orgID: %s", orgID)
	folder.PrettyPrint(orgFolder)

	// Example of retrieving all child folders
	fmt.Printf("\nRetrieving all child folders for 'stunning-horridus' under orgID %s:\n", orgID)
	childFolders := folderDriver.GetAllChildFolders(orgID, "stunning-horridus")
	folder.PrettyPrint(childFolders)

	// Uncomment the following code to test moving folders if needed
	/*
	name := "steady-insect"
	dst := "iruwhsdfniuwjke" // testing informally 

	fmt.Printf("\nMoving folder '%s' to new destination '%s'...\n", name, dst)
	updatedFolders, err := folderDriver.MoveFolder(name, dst)
	if err != nil {
		fmt.Printf("%s", err) // fix!
		return
	}
	folder.PrettyPrint(updatedFolders)
	*/
}