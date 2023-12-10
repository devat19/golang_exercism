package interest

func applicableInterestRate(bal float64) float32 {
	switch {
	case bal < 0:
		return 3.213
	case bal >= 0 && bal < 1000:
		return 0.5
	case bal >= 1000 && bal < 5000:
		return 1.621
	default:
		return 2.475
	}
}

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
	return applicableInterestRate(balance)
	// panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {

	var interestRate float32 = applicableInterestRate(balance)

	return (balance * float64(interestRate) / 100)
	// panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {

	return balance + Interest(balance)
	// panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {

	if targetBalance <= balance {
		return 0
	}

	var i int = 0
	for balance < targetBalance {
		i++
		balance = balance + Interest(balance)
	}
	return i
	// panic("Please implement the YearsBeforeDesiredBalance function")
}
