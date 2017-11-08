package tree

import "errors"

type Record struct {
	ID, Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

var children map[int][]int

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}
	// Sort children -> parent mapping and reject bad records
	// Based on the test, IDs must range from 0 to len(records)-1
	parents := make([]int, len(records))
	for _, r := range records {
		if r.ID < r.Parent || (r.ID != 0 && r.ID == r.Parent) || r.ID >= len(records) {
			return nil, errors.New("Error.")
		}
		parents[r.ID] = r.Parent
	}
	// Collect children of Nodes
	children = make(map[int][]int)
	for i := 1; i < len(parents); i++ {
		p := parents[i]
		children[p] = append(children[p], i)
	}
	return build(0), nil
}

// Build a Node recursively. Time complexity is O(n)
func build(id int) *Node {
	n := Node{ID: id}
	for _, i := range children[id] {
		n.Children = append(n.Children, build(i))
	}
	return &n
}
