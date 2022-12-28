package models

type Mechanical struct {
	NumberOfProfessors int
}

func (c *Mechanical) GetNumberOfProfessors() int {
	return c.NumberOfProfessors
}

func (c *Mechanical) GetName() string {
	return "mechanical"
}
