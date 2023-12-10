package isogram

import "unicode"

func IsIsogram(word string) bool {

	wordMap := make(map[rune]int)

	for _, letter := range word {

		if !unicode.IsLetter(letter) {
			continue
		}

		letter = unicode.ToLower(letter)

		_, exists := wordMap[letter]

		if exists {
			return false
		}

		// CHECK WITH LOGS THE EXACT RUNES HERE -> hpw to json stringify ??
		wordMap[letter] = 1

	}

	return true

	// forpanic("Please implement the IsIsogram function")
}
