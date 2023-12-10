package parsinglogfiles

import (
	"fmt"
	"regexp"
	"strings"
)

func IsValidLine(text string) bool {
	// panic("Please implement the IsValidLine function")
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`

	matched, err := regexp.MatchString(pattern, text)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	if matched {
		fmt.Println("[TRC] or [DBG] found at the start of the string!", text)
		return true
	}
	return false
}

func SplitLogLine(text string) []string {
	splitString, err := regexp.Compile(`\<[~*=-]*\>`)

	if err != nil {
		fmt.Println("Error:", err)
	}

	s := splitString.Split(text, -1)

	fmt.Println("[TRC] or [DBG] ", strings.Join(s, ""))

	return s
}

// func CountQuotedPasswords(lines []string) int {
// 		re := regexp.MustCompile(`(?i)\".*password.*\"`)
// 	for _, line := range lines {
// 		result += len(re.FindAllStringIndex(line, -1))
// 	}
// 	return

// count := 0
// 	r := regexp.MustCompile(`(?i)".*password.*"`)
// 	for _, line := range lines {
// 		if r.MatchString(line) {
// 			count++
// 		}
// 	}
// 	return count
// }

func CountQuotedPasswords(lines []string) (result int) {
	re := regexp.MustCompile(`(?i)\".*password.*\"`)
	for _, line := range lines {
		result += len(re.FindAllStringIndex(line, -1))
	}
	return
}

func RemoveEndOfLineText(text string) string {
	return regexp.MustCompile(`end-of-line\d+`).ReplaceAllLiteralString(text, "")
	// panic("Please implement the RemoveEndOfLineText function")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
	for i, line := range lines {
		match := re.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		lines[i] = line
	}
	return lines
	// panic("Please implement the TagWithUserName function")
}
