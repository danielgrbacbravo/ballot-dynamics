package main

import (
	"fmt"

	voting "danigrb.dev/ballot-dynamics/voter" // Update this with your actual module path
)

func main() {
	// Create a new FPTP voting system
	var votingSystem voting.VotingSystem
	votingSystem = voting.NewFPTPVotingSystem()

	var voters []voting.Voter
	voters = voting.GenerateRandomVoters(233)

	for _, voter := range voters {
		if err := votingSystem.RegisterVoter(voter); err != nil {
			fmt.Println("Error registering voter:", err)
		}
		fmt.Println("Registered voter:", voter.ID)
	}

	var candidates []voting.Candidate
	candidates = voting.GenerateRandomCandidates(2)

	// Cast votes
	for _, voter := range voters {
		if _, err := votingSystem.CastVote(voter, candidates); err != nil {
			fmt.Println("Error casting vote for voter:", err)

		}
	}

	// Report results
	results := votingSystem.ReportResults()
	fmt.Println(results)
}
