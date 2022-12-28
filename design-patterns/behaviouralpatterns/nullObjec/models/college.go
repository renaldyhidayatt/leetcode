package models

import "behaviouralpatterns/nullObjec/interfaces"

type College struct {
	Departments []interfaces.Department
}

func (c *College) AddDepartment(departmentName string, numOfProfessors int) {
	if departmentName == "computerscience" {
		computerScienceDepartment := &ComputerScience{NumberOfProfessors: numOfProfessors}
		c.Departments = append(c.Departments, computerScienceDepartment)
	}
	if departmentName == "mechanical" {
		mechanicalDepartment := &Mechanical{NumberOfProfessors: numOfProfessors}
		c.Departments = append(c.Departments, mechanicalDepartment)
	}

}

func (c *College) GetDepartment(departmentName string) interfaces.Department {
	for _, department := range c.Departments {
		if department.GetName() == departmentName {
			return department
		}
	}
	//Return a null department if the department doesn't exits
	return &NullDepartment{}
}
