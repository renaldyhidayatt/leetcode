package abstractfactory

import (
	"creationapatterns/abstractfactory/interfaces"
	"creationapatterns/abstractfactory/models"
	"fmt"
)

func AbstractFactory() {
	adidasFactory, _ := models.GetSportsFactory("adidas")

	nikeFactory, _ := models.GetSportsFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	nikeShort := nikeFactory.MakeShort()
	adidasShoe := adidasFactory.MakeShoe()
	adidasShort := adidasFactory.MakeShort()
	printShoeDetails(nikeShoe)
	printShortDetails(nikeShort)
	printShoeDetails(adidasShoe)
	printShortDetails(adidasShort)

}

func printShoeDetails(s interfaces.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}

func printShortDetails(s interfaces.IShort) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
