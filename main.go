package main

import (
	"fmt"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func main() {
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID) // converts organisation string ID into an UUID, storing it in orgID
	// returns nil if not valid UUID 

	res := folder.GetAllFolders() // calling get all folders from folder.go, part of folder package

	// example usage
	folderDriver := folder.NewDriver(res) // creates a new folder driver
	// creates a new struct interface, returned as IDriver interface type 
	// struct implements the interface 
	// interacting with struct through interface
	// struct satifies the interface - implements all methods (match signatures defined in interface)
	// get pointer to the struct, and Go knows it can call the interface methods 
	
	orgFolder := folderDriver.GetFoldersByOrgID(orgID) // filters based on organisation ID

	folder.PrettyPrint(res)
	fmt.Printf("\n Folders for orgID: %s", orgID) // prints all folders and filtered folders for specific ID
	folder.PrettyPrint(orgFolder) 

	// Calling GetAllChildFolders
	fmt.Printf("\nRetreiving all child folders for 'stunning-horridus' under orgID %s:\n", orgID)
	childFolders := folderDriver.GetAllChildFolders(orgID, "stunning-horridus")
	folder.PrettyPrint(childFolders)

	// Calling GetAllChildFolders (2)
	secondOrgId := uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7")
	fmt.Printf("\nRetrieving all child folders for 'creative-scalphunter' under orgID %s:\n", secondOrgId)
	childFoldersTwo := folderDriver.GetAllChildFolders(secondOrgId, "creative-scalphunter")
	folder.PrettyPrint(childFoldersTwo)

	// Initial folder structure
	fmt.Println("\n================= Initial Folder Structure =================")
	folder.PrettyPrint(res)

	// Move the folder
	name := "folder-name-to-move"     // Replace with the folder name you want to move
	dst := "new-destination-prefix"   // Replace with the desired new prefix

	fmt.Printf("\nMoving folder '%s' to new destination '%s'...\n", name, dst)
	updatedFolders, err := folderDriver.MoveFolder(name, dst)
	if err != nil {
		fmt.Println("Error moving folder:", err)
	} else {
		fmt.Println("Folder moved successfully! Updated folders:")
		folder.PrettyPrint(updatedFolders)
	}

	// Final folder structure
	fmt.Println("\n================= Updated Folder Structure =================")
	folder.PrettyPrint(folderDriver.GetAllFolders())



}

// creates a new instance of a driver, calling methods 
// folder.go defines driver struct (holds folder data) and IDriver interface
// get_folder.go defined the logic for the methods  