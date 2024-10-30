package folder_test

import (
	"testing"
    //"fmt"
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
	// "github.com/stretchr/testify/assert" // more expressive assertions
)

func Test_folder_GetFoldersByOrgID(t *testing.T) {
	t.Parallel()
    
    sampleFolders := folder.GetTestSampleData()

	tests := []struct {
        name    string
        orgID   uuid.UUID
        want    []folder.Folder
    }{
        {
            name:  "Test with valid orgID from sample data",
            orgID: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), // Updated orgID
            want: []folder.Folder{
                {
                    Name:  "capable-star-spangled",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled",
                },
                {
                    Name:  "golden-freefall",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall",
                },
                {
                    Name:  "devoted-burka",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka",
                },
                {
                    Name:  "definite-lizard",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.definite-lizard",
                },
                {
                    Name:  "moving-micromax",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.moving-micromax",
                },
                {
                    Name:  "trusted-superman",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman",
                },
                {
                    Name:  "stunning-black-panther",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman.stunning-black-panther",
                },
            },
        },
        {
            name:  "Test with valid orgID for ruling-michaelangelo",
            orgID: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), // Valid orgID from sample data
            want: []folder.Folder{
                {
                    Name: "ruling-michaelangelo",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo",
                },
                {
                    Name: "oriented-network",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.oriented-network",
                },
                {
                    Name: "dominant-wild",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.oriented-network.dominant-wild",
                },
                {
                    Name: "saving-beast",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.oriented-network.dominant-wild.saving-beast",
                },
                {
                    Name: "relative-siege",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.oriented-network.dominant-wild.relative-siege",
                },
                {
                    Name: "wanted-moonstar",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar",
                },
                {
                    Name: "sunny-warbound",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.sunny-warbound",
                },
                {
                    Name: "handy-mighty",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.sunny-warbound.handy-mighty",
                },
                {
                    Name: "correct-aquagirl",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.sunny-warbound.correct-aquagirl",
                },
                {
                    Name: "daring-firelord",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.daring-firelord",
                },
                {
                    Name: "grateful-madame-hydra",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.daring-firelord.grateful-madame-hydra",
                },
                {
                    Name: "free-raphael",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "ruling-michaelangelo.wanted-moonstar.daring-firelord.free-raphael",
                },
                {
                    Name: "curious-cable",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "curious-cable",
                },
                {
                    Name: "social-blastaar",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "curious-cable.social-blastaar",
                },
                {
                    Name: "keen-doorman",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "curious-cable.social-blastaar.keen-doorman",
                },
                {
                    Name: "allowed-the-punisher",
                    OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
                    Paths: "curious-cable.social-blastaar.keen-doorman.allowed-the-punisher",
                },
            },
        },
        {
            name:  "Test with valid, non-matching orgID using sample data",
            orgID: uuid.FromStringOrNil("12345678-1234-1234-1234-123456789abc"),
            want: []folder.Folder{}, // No folders expected
        },
        {
            name:  "Test with zero-value (invalid) orgID using sample data",
            orgID: uuid.Nil,
            want: []folder.Folder{}, // No folders expected
        },
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := folder.NewDriver(sampleFolders)
            got := f.GetFoldersByOrgID(tt.orgID)

            if len(got) != len(tt.want) {
                t.Errorf("expected length %d, got %d", len(tt.want), len(got))
            }

            for i := range tt.want {
                if got[i] != tt.want[i] {
                    t.Errorf("expected folder %v, got %v", tt.want[i], got[i])
                }
            }
        })
    }
}

// Test for getting all child folders
func Test_folder_GetAllChildFolders(t *testing.T) {
	t.Parallel()

	// Load sample data
	sampleFolders := folder.GetTestSampleData()

	tests := []struct {
		name       string // name of test
		orgID      uuid.UUID
		parentName string // parent folder name (to search for)
		want       []folder.Folder
	}{
		{
			name: "Test for parent folder 'capable-star-spangled'",
			orgID: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
			parentName: "capable-star-spangled",
			want: []folder.Folder{
                {
                    Name: "golden-freefall", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall"},
                {
                    Name: "devoted-burka", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka"},
                {
                    Name: "definite-lizard", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.definite-lizard"},
                {
                    Name: "moving-micromax", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.moving-micromax"},
                {
                    Name: "trusted-superman", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman"},
                {
                    Name: "stunning-black-panther", 
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"), 
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman.stunning-black-panther"},
            },
		},
		{
			name: "Test for parent folder 'golden-freefall'",
			orgID: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
			parentName: "golden-freefall",
			want: []folder.Folder{
                {
                    Name: "devoted-burka",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka",
                },
                {
                    Name: "definite-lizard",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.definite-lizard",
                },
                {
                    Name: "moving-micromax",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.devoted-burka.moving-micromax",
                },
                {
                    Name: "trusted-superman",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman",
                },
                {
                    Name: "stunning-black-panther",
                    OrgId: uuid.FromStringOrNil("e546a29e-4432-4243-8e71-f26aa8fbddbd"),
                    Paths: "capable-star-spangled.golden-freefall.trusted-superman.stunning-black-panther",
                },
            },
		},
		{
			name: "Test with non-matching orgID",
			orgID: uuid.FromStringOrNil("12345678-1234-1234-1234-123456789abc"),
			parentName: "capable-star-spangled",
			want: []folder.Folder{}, // No folders expected
		},
		{
			name: "Test with zero-value (invalid) orgID",
			orgID: uuid.Nil,
			parentName: "capable-star-spangled",
			want: []folder.Folder{}, // No folders expected
		},
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            f := folder.NewDriver(sampleFolders) // Creates a driver with the sample data
            got := f.GetAllChildFolders(tt.orgID, tt.parentName) // Calls the function being tested

            if len(got) != len(tt.want) {
                t.Errorf("expected length %d, got %d", len(tt.want), len(got))
            }

            for i := range tt.want {
                if got[i] != tt.want[i] {
                    t.Errorf("expected folder %v, got %v", tt.want[i], got[i])
                }
            }
        })
    }
}
