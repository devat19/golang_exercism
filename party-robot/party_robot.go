package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	// panic("Please implement the Welcome function")
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	// panic("Please implement the HappyBirthday function")
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// functions with first letter as small case aren't exported
// creating a function to add zeroes to a string
// %03d  < This would work the same
func padLeft(str string, length int) string {
	for len(str) < length {
		str = "0" + str
	}
	return str
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// panic("Please implement the AssignTable function")
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.
	return fmt.Sprintf(Welcome(name) + "\n" +
		fmt.Sprintf("You have been assigned to table %v. Your table is %s, exactly %.1f meters from here.", padLeft(fmt.Sprintf("%v", table), 3), direction, distance) +
		"\n" + fmt.Sprintf("You will be sitting next to %s.", neighbor))
}
