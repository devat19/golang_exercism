package booking

import "time"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// panic("Please implement the Schedule function")
	t, err := time.Parse("1/02/2006 15:04:05", date) // Practical tip - understand worldwide time formatting system
	if err != nil {                                  // Alwa-ys check errors even if they should not happen.
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {

	formattedDate, err := time.Parse("January 2, 2006 15:04:05", date)
	if err != nil { // Alwa-ys check errors even if they should not happen.
		panic(err)
	}

	return time.Now().After(formattedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	formattedDate, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil { // Alwa-ys check errors even if they should not happen.
		panic(err)
	}

	var hour int = formattedDate.Hour()

	// if(hour >= 12 && hour <= 18) {
	// 	return true
	// }
	return hour >= 12 && hour <= 18
	// panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	formattedDate, err := time.Parse("1/2/2006 15:04:05", date)
	if err != nil { // Alwa-ys check errors even if they should not happen.
		panic(err)
	}

	return formattedDate.Format("You have an appointment on Monday, January 2, 2006, at 15:04.")

	// panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	formattedDate, err := time.Parse("1/2/2006 15:04:05", "9/15/2023 00:00:00")
	if err != nil { // Alwa-ys check errors even if they should not happen.
		panic(err)
	}
	return formattedDate

	//     return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

	// panic("Please implement the AnniversaryDate function")
}
