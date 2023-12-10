package chessboard

// CHECK generics

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File = []bool
type Chessboard = map[string]File

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {

	var count int = 0
	for k, v := range cb {
		if k == file {
			for _, val := range v {
				if val {
					count++
				}
			}
		}
	}
	return count
	// panic("Please implement CountInFile()")
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {

	var count int = 0
	for _, v := range cb {

		for idx, val := range v {
			if idx == rank-1 && val {
				count++
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {

	var count int = 0
	for _, v := range cb {

		count = count + len(v)
	}
	return count
	// panic("Please implement CountAll()")
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {

	var count int = 0
	for _, v := range cb {

		for _, val := range v {
			if val {
				count++
			}
		}
	}
	return count
	// panic("Please implement CountOccupied()")
}
