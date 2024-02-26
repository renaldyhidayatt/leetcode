package partyrobot

import "fmt"

func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
}

func AssignTable(guestName string, tableNumber int, seatmate string, direction string, distance float64) string {
	tableNumberStr := fmt.Sprintf("%03d", tableNumber)

	distanceFormatted := fmt.Sprintf("%.1f", distance)

	return fmt.Sprintf("%s\nYou have been assigned to table %s. Your table is %s, exactly %s meters from here.\nYou will be sitting next to %s.", Welcome(guestName), tableNumberStr, direction, distanceFormatted, seatmate)
}
