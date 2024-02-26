package techpalace

import "strings"

func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

func AddBorder(welcomeMsg string, numStartsPerline int) string {
	border := strings.Repeat("*", numStartsPerline)

	return border + "\n" + welcomeMsg + "\n" + border
}

func CleanupMessage(oldMsg string) string {
	cleanedMsg := strings.ReplaceAll(oldMsg, "*", "")

	cleanedMsg = strings.TrimSpace(cleanedMsg)

	return cleanedMsg
}
