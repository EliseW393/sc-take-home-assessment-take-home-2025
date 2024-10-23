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
}

// creates a new instance of a driver, calling methods 
// folder.go defines driver struct (holds folder data) and IDriver interface
// get_folder.go defined the logic for the methods  