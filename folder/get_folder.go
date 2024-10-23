package folder

import "github.com/gofrs/uuid"
import (
	"strings" // string manipulation functions
	"encoding/json" // working with JSON
)

// driver struct not required 
func GetAllFolders() []Folder { // returns all folders 
	return GetSampleData()
}

// two functions below require a driver struct
// (f *driver) the reciever: f is the reciever variable, refers to current instance of driver struct
// the * means f is a pointer to the driver struct
func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder { // returns all folders by specific orgID (loop to filter)
	folders := f.folders // accessing folders in driver struct

	res := []Folder{}
	for _, f := range folders { // f the loop variable representing each individual Folder struct, underscore ignores index
		if f.OrgId == orgID { // checks if org ID of current Folder object matches parameter org ID
			res = append(res, f) // adding to the res slice (dynamic array)
		}
	}

	return res

}

// method belongs to the driver struct 
// working with a pointer to the driver struct - memory, performance
// assumed driver struct exists, f is its pointer
// has attribute with slide of all Folders read from JSON file
// 


//[]Folder - a slice of type Folder
//{} empty
// only one type, unless you declare slice of interfaces 
// interface defines a set of methods a type must implement 
// type that implements all those methods satisfies the interface 
// but can have different behaviours for those methods 
func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder { // don't forget function description!
	
// Create the []Folder{} slice
	// Retrieve all folders for the given orgID using GetFoldersByOrgID, put in variable
	
	// Create the []Folder{} slice

	// iterate through filtered folders, check for parent name -> if yes, append to Folder

	// return Folders - slice of Folder objects (structs)

	return []Folder{}
}




// implementation for fetching folders baded on certain criteria

// list of paths represented as strings in JSON file, dot separated 

// notes
// include: all child folders, including deeper nested ones
// e.g. if given folder is 'b': a.b.c and d.b.c included, a.b.c and a.b.f included 

// if something like alpha.bravo.charlie, searching for bravo, would include charlie, with this path : "alpha.bravo.charlie"


// type Folder struct {
//    Name  string
//    OrgID uuid.UUID
//    Path  string
//}
