package uber

import (
	"reflect"
	"testing"
)

func TestPrintDirectoryIndentation(t *testing.T) {
	root := &DirNode{
		Name: "dir",
		Children: []*DirNode{
			{
				Name: "subdir1",
				Children: []*DirNode{
					{Name: "file1.ext"},
					{Name: "subsubdir1"},
				},
			},
			{
				Name: "subdir2",
				Children: []*DirNode{
					{
						Name: "subsubdir2",
						Children: []*DirNode{
							{Name: "file2.ext"},
						},
					},
				},
			},
		},
	}

	got := PrintDirectory(root)
	want := []string{
		"dir",
		" subdir1",
		"  file1.ext",
		"  subsubdir1",
		" subdir2",
		"  subsubdir2",
		"   file2.ext",
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("PrintDirectory() = %v, want %v", got, want)
	}
}
