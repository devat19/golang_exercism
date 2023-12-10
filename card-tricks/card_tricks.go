package cards

import "fmt"

/*
  Print bool value
  What's main package? How to use debugger alongwith and how to use debugger on individual files
*/

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2, 6, 9}
	// panic("Please implement the FavoriteCards function")
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
// Add error handling -> return that way - check stack trace, and return from catch that way
func GetItem(slice []int, index int) int {
	fmt.Printf("%v, %v, %v", len(slice), index, slice)
	if len(slice) <= index {
		return -1
	}
	if index < 0 {
		return -1
	}
	return slice[index]
	// panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
// Can slices have different data types? Will this explode if the data type of value wasn't int
// Add error handling
func SetItem(slice []int, index, value int) []int {
	if index >= len(slice) {
		return append(slice, value)
	} else {
		if index < 0 {
			slice = append(slice, value)
		} else {
			slice[index] = value
		}
	}
	return slice
	// panic("Please implement the SetItem function")
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
	// panic("Please implement the PrependItems function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {

	var sl []int

	for idx, value := range slice {
		if index != idx {
			sl = append(sl, value)
		}
	}
	return sl

	/*
		if index < 0 || index >= len(slice) {
			return slice
		}
		return append(slice[:index], slice[index+1:]...)
	*/

	// panic("Please implement the RemoveItem function")
}
