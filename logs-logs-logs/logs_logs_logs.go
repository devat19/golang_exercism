package logs

import (
	"fmt"
	"unicode/utf8"
)

func indentifyApplication(unicode rune) string {
	var unicodeVal string = fmt.Sprintf("%U", unicode)
	switch {
	case unicodeVal == "U+2757":
		return "recommendation"
	case unicodeVal == "U+1F50D":
		return "search"
	case unicodeVal == "U+2600":
		return "weather"
	default:
		return "default"
	}
}

// Application identifies the application emitting the given log.
func Application(log string) string {

	var application string = "default"

	for _, char := range log {
		var app string = indentifyApplication(char)

		if app == "recommendation" || app == "search" || app == "weather" {
			application = app
			break
		}
	}

	return application
	// panic("Please implement the Application() function")
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {

	var repalcedString string = ""

	// char is a rune
	for _, char := range log {
		if char == oldRune {
			repalcedString = repalcedString + string(newRune)
		} else {
			repalcedString = repalcedString + string(char)
		}
	}

	return repalcedString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {

	return utf8.RuneCountInString(log) <= limit
	// return len([]rune(log)) <= limit  << This also works!
	// panic("Please implement the WithinLimit() function")
}
