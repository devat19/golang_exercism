package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totBirds int = 0

	for _, birdCount := range birdsPerDay {
		totBirds = totBirds + birdCount
	}
	return totBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var sumBirdsInWeek int = 0

	var daysIndexSlice []int = birdsPerDay[(week-1)*7 : (week-1)*7+7]

	for i := 0; i < len(daysIndexSlice); i++ {
		sumBirdsInWeek = sumBirdsInWeek + daysIndexSlice[i]
	}
	return sumBirdsInWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			birdsPerDay[i] = birdsPerDay[i] + 1
		}
	}

	return birdsPerDay
	// panic("Please implement the FixBirdCountLog() function")
}
