package nullobjec

import (
	"behaviouralpatterns/nullObjec/models"
	"fmt"
)

func NullObject() {

	college1 := createCollege1()
	college2 := createCollege2()
	totalProfessors := 0
	departmentArray := []string{"computerscience", "mechanical", "civil", "electronics"}
	for _, deparmentName := range departmentArray {
		d := college1.GetDepartment(deparmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college1 is %d\n", totalProfessors)
	totalProfessors = 0
	for _, deparmentName := range departmentArray {
		d := college2.GetDepartment(deparmentName)
		totalProfessors += d.GetNumberOfProfessors()
	}
	fmt.Printf("Total number of professors in college2 is %d\n", totalProfessors)
}

func createCollege1() *models.College {
	college := &models.College{}
	college.AddDepartment("computerscience", 4)
	college.AddDepartment("mechanical", 5)
	return college
}

func createCollege2() *models.College {
	college := &models.College{}
	college.AddDepartment("computerscience", 2)
	return college
}
