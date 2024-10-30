package folder_test

import (
	"testing"
	"errors"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_MoveFolder(t *testing.T) {
	t.Parallel()

	originalSampleFolders := folder.GetTestSampleData()

	tests := []struct {
		name           string
		source         string
		dst            string
		expectedResult []folder.Folder
		expectedErr    error
	}{
		{
			name:   "Move folder 'wanted-moonstar' under 'oriented-network'",
			source: "wanted-moonstar",
			dst:    "oriented-network",
			expectedResult: []folder.Folder{
			{
				Name: "wanted-moonstar",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar",
			},
			{
				Name: "sunny-warbound",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.sunny-warbound",
			},
			{
				Name: "handy-mighty",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.sunny-warbound.handy-mighty",
			},
			{
				Name: "correct-aquagirl",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.sunny-warbound.correct-aquagirl",
			},
			{
				Name: "daring-firelord",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.daring-firelord",
			},
			{
				Name: "grateful-madame-hydra",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.daring-firelord.grateful-madame-hydra",
			},
			{
				Name: "free-raphael",
				OrgId: uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"),
				Paths: "ruling-michaelangelo.oriented-network.wanted-moonstar.daring-firelord.free-raphael",
			},
		},
			expectedErr: nil,
	},
		{
			name:   "Attempt to move folder to itself",
			source: "trusted-superman",
			dst:    "trusted-superman",
			expectedResult: originalSampleFolders,
			expectedErr: errors.New("Error: Cannot move a folder to itself\n"),
		},

		// Test for 'child of itself' condition: known issue where test fails despite successful manual testing.
		{ 
			name:   "Move folder to a child of itself",
			source: "capable-star-spangled",
			dst:    "golden-freefall",
			expectedResult: originalSampleFolders,
			expectedErr: errors.New("Error: Cannot move a folder to a child of itself\n"),
		},
		{
			name:   "Move folder to a different organisation",
			source: "capable-star-spangled",
			dst:    "ruling-michaelangelo",
			expectedResult: originalSampleFolders,
			expectedErr: errors.New("Error: Cannot move a folder to a different organisation\n"),
		},
		{
			name:   "Move folder with empty destination",
			source: "capable-star-spangled",
			dst:    "",
			expectedResult: originalSampleFolders,
			expectedErr: errors.New("Error: Source and destination folder names must not be empty\n"),
		},
		{
			name:   "Move folder to a parent of itself",
			source: "golden-freefall",
			dst:    "capable-star-spangled",
			expectedResult: originalSampleFolders,
			expectedErr: errors.New("Error: Destination folder is already a parent of the source folder, cannot move.\n"),
		},
	}

	for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {

            // deep copy of original sample data for isolation
            sampleFolders := make([]folder.Folder, len(originalSampleFolders))
            copy(sampleFolders, originalSampleFolders)
            
            f := folder.NewDriver(sampleFolders) // driver with fresh data for each test
            got, err := f.MoveFolder(tt.source, tt.dst)

            if tt.expectedErr != nil {
                if err == nil || err.Error() != tt.expectedErr.Error() {
                    t.Errorf("expected error %v, got %v", tt.expectedErr, err)
                }
            } else if err != nil {
                t.Errorf("unexpected error: %v", err)
            }

            // comparing only the folders that are expected to change
            for _, expectedFolder := range tt.expectedResult {
                found := false
                for _, resultFolder := range got {
                    if resultFolder.Name == expectedFolder.Name {
                        if resultFolder.Paths != expectedFolder.Paths {
                            t.Errorf("for folder %s: expected path %s, got %s", resultFolder.Name, expectedFolder.Paths, resultFolder.Paths)
                        }
                        found = true
                        break
                    }
                }
                if !found {
                    t.Errorf("expected folder %s not found in result", expectedFolder.Name)
                }
            }
        })
    }
}