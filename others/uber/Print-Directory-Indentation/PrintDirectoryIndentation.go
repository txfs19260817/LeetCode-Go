package uber

import "strings"

// DirNode represents a directory tree node (file or directory).
type DirNode struct {
	Name     string
	Children []*DirNode
}

// PrintDirectory returns the directory listing as lines with indentation.
// It performs a preorder traversal and prints children in input order.
func PrintDirectory(root *DirNode) []string {
	var lines []string
	var dfs func(node *DirNode, depth int)
	dfs = func(node *DirNode, depth int) {
		if node == nil {
			return
		}
		lines = append(lines, strings.Repeat(" ", depth)+node.Name)

		for _, child := range node.Children {
			dfs(child, depth+1)
		}
	}

	dfs(root, 0)
	return lines
}
