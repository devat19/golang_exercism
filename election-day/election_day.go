package electionday

import "fmt"

/*

var a int
a = 2

var p *int
p = &a // the variable 'p' contains the memory address of 'a'

var b int
b = *p // b == 2  dereferencing

*/
// NewVoteCounter returns a new vote counter with
// a given number of initial votes.
func NewVoteCounter(initialVotes int) *int {

	fmt.Println(&initialVotes)
	return &initialVotes
	// panic("Please implement the NewVoteCounter() function")
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {

	if counter == nil {
		return 0
	}
	fmt.Println(*counter)
	return *counter
	// panic("Please implement the VoteCount() function")
}

// IncrementVoteCount increments the value in a vote counter.
func IncrementVoteCount(counter *int, increment int) {
	*counter = *counter + increment
	// panic("Please implement the IncrementVoteCount() function")
}

// NewElectionResult creates a new election result.
func NewElectionResult(candidateName string, votes int) *ElectionResult {

	return &ElectionResult{Name: candidateName, Votes: votes}
	// panic("Please implement the NewElectionResult() function")
}

// DisplayResult creates a message with the result to be displayed.
func DisplayResult(result *ElectionResult) string {
	return fmt.Sprintf("%s (%v)", result.Name, result.Votes)
	// panic("Please implement the DisplayResult() function")
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map.
func DecrementVotesOfCandidate(results map[string]int, candidate string) {

	candidateVote, candidateExists := results[candidate]

	if !candidateExists {
		return
	}

	results[candidate] = candidateVote - 1
	// panic("Please implement the DecrementVotesOfCandidate() function")
}
