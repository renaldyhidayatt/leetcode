package models

type NullDepartment struct {
	NumberOfProfessors int
}

func (c *NullDepartment) GetNumberOfProfessors() int {
	return 0
}

func (c *NullDepartment) GetName() string {
	return "null department"
}
