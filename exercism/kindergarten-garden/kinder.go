package kindergartengarden

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Garden struct {
	plants   []string
	students []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	if hasInvalidCodes(diagram) {
		return nil, errors.New("error: invalid cup codes")
	}

	if hasDuplicateChild(children) {
		return nil, errors.New(fmt.Sprintf("error: the student odd student list"))
	}

	if !strings.HasPrefix(diagram, "\n") || strings.Count(diagram, "\n") != 2 {
		return nil, errors.New("Invalid diagram")
	}

	plantMap := map[string]string{
		"V": "violets",
		"G": "grass",
		"C": "clover",
		"R": "radishes",
	}

	diagram = strings.Replace(diagram, "\n", "", 1)

	rows := strings.Split(diagram, "\n")
	rows0 := strings.Split(rows[0], "")
	rows1 := strings.Split(rows[1], "")

	if len(children)*4 != (len(diagram) - 1) {
		return nil, errors.New("error: not enough plant or student")
	}

	garden := Garden{
		plants:   []string{},
		students: children,
	}

	for i, _ := range children {
		garden.plants = append(garden.plants, plantMap[rows0[i*2]])
		garden.plants = append(garden.plants, plantMap[rows0[i*2+1]])
		garden.plants = append(garden.plants, plantMap[rows1[i*2]])
		garden.plants = append(garden.plants, plantMap[rows1[i*2+1]])
	}
	return &garden, nil

}

func (g *Garden) Plants(child string) ([]string, bool) {
	sortedStudents := g.SortedStudents()

	for i, v := range sortedStudents {
		if child == v {
			return g.plants[i*4 : i*4+4], true
		}
	}

	return nil, false
}

func (g *Garden) SortedStudents() []string {
	newSlc := []string{}

	for _, v := range g.students {
		newSlc = append(newSlc, v)
	}

	sort.Strings(newSlc)

	return newSlc
}

func hasInvalidCodes(diagram string) bool {
	diagram = strings.ReplaceAll(diagram, "\n", "")

	for _, v := range diagram {
		if v != 'V' && v != 'C' && v != 'G' && v != 'R' {
			return true
		}
	}

	return false
}

func hasDuplicateChild(children []string) bool {
	childrenStr := strings.Join(children, ",")

	for _, v := range children {
		if v == "" {
			return true
		}

		if strings.Count(childrenStr, v) > 1 {
			return true
		}
	}

	return false
}
