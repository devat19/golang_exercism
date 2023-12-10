package techpalace

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	capitalizedProper := fmt.Sprintf("Welcome to the Tech Palace, %s", cases.Title(language.English, cases.Compact).String(customer))
	fmt.Println((capitalizedProper))
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToTitle(customer))
	// panic("Please implement the WelcomeMessage() function")
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
	return strings.Repeat("*", numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat("*", numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
// REDO -
// https://stackoverflow.com/questions/1103985/method-chaining-why-is-it-a-good-practice-or-not
// can we do this without constant assign to spacedString ?
func CleanupMessage(oldMsg string) string {
	var spacedString string = strings.ReplaceAll(oldMsg, "*", " ")
	spacedString = strings.Trim(spacedString, " ")
	spacedString = strings.Trim(spacedString, "\n")
	return strings.Trim(spacedString, " ")
}
