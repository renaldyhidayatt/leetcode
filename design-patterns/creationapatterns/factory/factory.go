package factory

import (
	getgun "creationapatterns/factory/getGun"
	"creationapatterns/factory/interfaces"
	"fmt"
)

func Factory() {
	ak47, _ := getgun.GetGun("ak47")
	maverick, _ := getgun.GetGun("maverick")
	printDetails(ak47)
	printDetails(maverick)

}

func printDetails(g interfaces.IGun) {
	fmt.Printf("Gun: %s", g.GetName())
	fmt.Println()
	fmt.Printf("Power: %d", g.GetPower())
	fmt.Println()
}
