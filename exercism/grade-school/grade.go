package gradeschool

import "sort"

type Grade struct {
	level    int
	students []string
}

type School map[int][]string

func New() *School {
	return &School{}
}

func (s School) Add(student string, grade int) {
	s[grade] = append(s[grade], student)
}

func (s School) Grade(level int) []string {
	return s[level]
}

func (s School) Enrollment() []Grade {
	var rooster []Grade

	for level, student := range s {
		sort.Strings(student)

		rooster = append(rooster, Grade{level, student})
	}

	sort.Slice(rooster, func(i, j int) bool {
		return rooster[i].level < rooster[j].level
	})

	return rooster
}
