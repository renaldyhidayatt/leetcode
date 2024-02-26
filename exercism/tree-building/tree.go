package treebuilding

import "fmt"

const rootID = 0

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) == 0 {
		return nil, nil
	}

	positions := make([]int, len(records))

	for i, r := range records {
		if r.ID < rootID || r.ID >= len(records) {
			return nil, fmt.Errorf("out of bounds record id %d", r.ID)
		}
		positions[r.ID] = i
	}

	nodes := make([]Node, len(records))

	for i := range positions {
		r := records[positions[i]]

		if r.ID != i {
			return nil, fmt.Errorf("non-contiguous node %d (want %d)", r.ID, i)
		}

		validParentChild := (r.ID > r.Parent) || (r.ID == rootID && r.Parent == rootID)

		if !validParentChild {
			return nil, fmt.Errorf("node %d hash impossible parent %d", r.ID, r.Parent)
		}

		nodes[i].ID = i

		if i != rootID {
			p := &nodes[r.Parent]
			p.Children = append(p.Children, &nodes[i])
		}
	}

	return &nodes[0], nil
}
