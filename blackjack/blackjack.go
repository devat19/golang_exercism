package blackjack

func cardValue(card string) int {
	switch card {
	case "ace":
		return 11
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "one":
		return 1
	default:
		return 0
	}
}

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	return cardValue(card)
	// panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {

	var card1Value int = cardValue(card1)
	var card2Value int = cardValue(card2)
	var dealerCardValue int = cardValue(dealerCard)

	var sum int = card1Value + card2Value

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case sum == 21 && !(dealerCard == "ace" || dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king"):
		return "W"
	case sum == 21 && (dealerCard == "ace" || dealerCard == "jack" || dealerCard == "queen" || dealerCard == "king"):
		return "S"
	case sum >= 17 && sum <= 20:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardValue < 7:
		return "S"
	case sum >= 12 && sum <= 16 && dealerCardValue >= 7:
		return "H"
	default:
		return "H"
	}
}
