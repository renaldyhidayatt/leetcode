package models

type ComputerScience struct {
	NumberOfProfessors int
}

func (c *ComputerScience) GetNumberOfProfessors() int {
	return c.NumberOfProfessors

}

func (c *ComputerScience) GetName() string {
	return "computerscience"
}
