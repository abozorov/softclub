package main

import "fmt"

func main() {
	votes := []string{"Ali", "Umed", "Ali", "Azam", "Ali", "Umed"}
	voteCount := make(map[string]int, len(votes))
	mxVote := 0

	for _, v := range votes {
		voteCount[v] ++
		if mxVote < voteCount[v] {
			mxVote = voteCount[v]
		}
	}

	top := make(map[string]struct{}, len(votes))
	for _, v := range votes {
		if mxVote == voteCount[v] {
			top[v] =struct{}{}
		}
	}

	fmt.Printf("%v\n", top)
}
