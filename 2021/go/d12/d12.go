package d12

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	name     string
	small    bool
	children *[]Node
}

func createNode(name string) (n Node) {
	n.name = name
	n.children = &([]Node{})
	n.small = !(strings.ToUpper(name) == name)
	return
}

func getData(path string) (root Node) {
	nodes := make(map[string]Node)

	f, _ := os.Open(path)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()

		splits := strings.Split(text, "-")

		n1 := splits[0]
		_, ok := nodes[n1]
		if !ok {
			nodes[n1] = createNode(n1)
		}

		n2 := splits[1]
		_, ok = nodes[n2]
		if !ok {
			nodes[n2] = createNode(n2)
		}

		c := nodes[n1].children
		*c = append(*c, nodes[n2])

		c = nodes[n2].children
		*c = append(*c, nodes[n1])
	}

	// some cleanup
	// very hacky, I could make that cleaner but OK for this
	x := nodes["end"].children
	*x = []Node{}

	return nodes["start"]
}

func isPathValid(path []Node, part int) bool {
	/*  part 2 is:
	a single small cave can be visited at most twice,
	and the remaining small caves can be visited at most once
	*/
	small := make(map[string]int)
	switch part {
	case 1:
		for _, node := range path {
			if node.small {
				if _, ok := small[node.name]; !ok {
					small[node.name] = 1
				} else {
					return false
				}
			}
		}
	case 2:
		twiceFlag := false
		for _, node := range path {
			if node.small {
				if _, ok := small[node.name]; !ok {
					small[node.name] = 1
				} else if !twiceFlag && node.name != "start" {
					small[node.name] += 1
					twiceFlag = true
				} else {
					return false
				}
			}
		}
	}
	return true
}

// Depth-first Search (DFS)
// if the node is a leaf, print the list/stack
// pop the node from the list/stack

func search(root Node, paths *[][]Node, part int) {
	n := len(*paths)
	(*paths)[n-1] = append((*paths)[n-1], root)

	if isPathValid((*paths)[n-1], part) {
		for _, n := range *root.children {
			search(n, paths, part)
		}
	}

	n = len(*paths)
	path := (*paths)[n-1]
	newPath := make([]Node, len(path)-1)
	copy(newPath, path[:len(path)-1])
	if len(*root.children) == 0 && path[len(path)-1].name == "end" {
		// leaf
		(*paths) = append((*paths), newPath)
	} else {
		(*paths)[n-1] = newPath
	}
}

func Solve(path string, part int) {
	root := getData(path)
	var paths [][]Node
	paths = append(paths, []Node{})
	search(root, &paths, part)

	// process paths
	nPaths := 0
	for _, path := range paths {
		if len(path) > 0 {
			text := ""
			for _, n := range path {
				text += fmt.Sprintf("%s,", n.name)
			}
			fmt.Println(text[:len(text)-1])
			nPaths++
		}
	}

	fmt.Printf("Number of paths: %d\n", nPaths)
}
