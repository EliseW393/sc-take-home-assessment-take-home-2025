package folder_test

import (
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert" // more expressive assertions
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()

	// each struct a test case
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
		// call GetSampleData for given Folders

		// this is a single test case with own created folders
		{
            name:  "Test with valid orgID",
            orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), // valid orgID
            folders: []folder.Folder{
                {Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
                {Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"},
            },
            want: []folder.Folder{
                {Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"},
                {Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"},
            },
        }

		// can add another test case here

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) { // tt: single test case from the slice, tt.folders is folders field of current test slice, gets single test case and assigns it tt
			f := folder.NewDriver(tt.folders) // driver with test folders

			get := f.GetFoldersByOrgID(tt.orgID) // calling the function being tested 

			// common assertions 
			// reflect.DeepEqual - basic equality check; t.Errorf("Function() = %v, want %v", got, want)
			// got != nill; t.Errorf("Expected nil, got %v", got)
			// len(got) != expectedLength; t.Errorf("Expected length %d, got %d", expectedLength, len(got))
			// err != nill; t.Errorf("Expected no error, got %v", err)
			// got != want; t.Errorf("Expected %q, got %q", want, got)

		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name    string // name of test 
        orgID   uuid.UUID
        parentName    string // parent folder name (to search for)
        folders []folder.Folder
        want    []folder.Folder 
    }{
        {
            name:  "Test with valid parent folder",
            orgID: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"),
            parentName:  "bravo", // The parent folder to search for
            folders: []folder.Folder{
                {Name: "bravo", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo"}, // no children 
                {Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"}, // charlie as child 
                {Name: "delta", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.delta"}, // not present
            },
            want: []folder.Folder{
                {Name: "charlie", OrgId: uuid.FromStringOrNil("38b9879b-f73b-4b0e-b9d9-4fc4c23643a7"), Paths: "alpha.bravo.charlie"}, // extra comma 
            },
        }
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := folder.NewDriver(tt.folders) // Create a driver with test folders
            got := f.GetAllChildFolders(tt.orgID, tt.name) // Call the function being tested

			// add asserts here 

        })
    }
}

package folder_test

import (
    "testing"
    "github.com/georgechieng-sc/interns-2022/folder"
    "reflect"
)

func Test_folder_GetUniqueChildren(t *testing.T) {
    t.Parallel()

    tests := []struct {
        name          string    
        existing      []string  // existing list of children
        path          string    // path with including names to add
        expected      []string  // expected output after adding unique names
    }{
        {
            name:     "Test adding multiple unique children",
            existing: []string{"charlie"}, 
            path:     "alpha.bravo.charlie.delta", 
            expected: []string{"charlie", "alpha", "bravo", "delta"}, 
        },
        {
            name:     "Test with no existing children",
            existing: []string{}, 
            path:     "alpha.bravo.delta", 
            expected: []string{"alpha", "bravo", "delta"}, 
        },
        {
            name:     "Test with duplicates",
            existing: []string{"charlie"}, 
            path:     "alpha.bravo.charlie", 
            expected: []string{"charlie", "alpha", "bravo"}, 
        },
        {
            name:     "Test with multiple unique children to empty list",
            existing: []string{}, 
            path:     "alpha.bravo.charlie.delta.epsilon", 
            expected: []string{"alpha", "bravo", "charlie", "delta", "epsilon"}, 
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            
			var got []string // hold unique children

            f := folder.NewDriver(nil) // (assuming no initial folders needed)

            for _, child := range tt.existing { // unique children slice with values from "existing"
                got = append(got, child)
            }

            f.GetUniqueChildren(tt.path, &got) // calling the helper function, passing the address of the slice

            // Assert that the result from the function matches the output

        })
    }
}



