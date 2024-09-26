package voting

import (
	"errors"
	"fmt"
)

// FPTPVotingSystem struct to implement VotingSystem
type FPTPVotingSystem struct {
	voters     map[int]Voter
	candidates map[string]Candidate
	votes      []Vote
}

// NewFPTPVotingSystem constructor
func NewFPTPVotingSystem() *FPTPVotingSystem {
	return &FPTPVotingSystem{
		voters:     make(map[int]Voter),
		candidates: make(map[string]Candidate),
		votes:      []Vote{},
	}
}

//TODO: Implement the RegisterVoter, CastVote, CountVotes, and ReportResults methods for the FPTP voting system

// RegisterVoter implements the RegisterVoter method
func (f *FPTPVotingSystem) RegisterVoter(voter Voter) error {
	if _, exists := f.voters[voter.ID]; exists {
		return errors.New("voter already registered")
	}
	f.voters[voter.ID] = voter
	return nil
}

//BUG: The CastVote method is not implemented correctly
// i need to build a function that will take the voter and the candidates and return the vote

// CastVote implements the CastVote method
func (f *FPTPVotingSystem) CastVote(voter Voter, candidates []Candidate) (Vote, error) {
	if _, exists := f.voters[voter.ID]; !exists {
		return Vote{}, errors.New("voter not registered")
	}

	// Determine which candidate the voter prefers based on their preferences
	var selectedCandidateID int
	maxScore := -1.0
	for _, candidate := range candidates {
		if score, exists := voter.Preferences[candidate.ID]; exists && score > maxScore {
			maxScore = score
			selectedCandidateID = candidate.ID
		}
	}

	if selectedCandidateID == "" {
		return Vote{}, errors.New("no valid candidate selected")
	}

	vote := Vote{
		VoterID:     voter.ID,
		CandidateID: selectedCandidateID,
		Score:       maxScore, // Use the maximum score as the vote score
	}
	f.votes = append(f.votes, vote)
	return vote, nil
}

// CountVotes implements the CountVotes method for FPTP
func (f *FPTPVotingSystem) CountVotes(votes []Vote) (map[string]int, error) {
	results := make(map[string]int)
	for _, vote := range votes {
		results[vote.CandidateID]++
	}
	return results, nil
}

// ReportResults implements the ReportResults method
func (f *FPTPVotingSystem) ReportResults() string {
	counts, _ := f.CountVotes(f.votes)
	var winnerID string
	maxVotes := 0

	// Determine the candidate with the most votes
	for candidate, count := range counts {
		if count > maxVotes {
			maxVotes = count
			winnerID = candidate
		}
	}

	resultStr := "Election Results:\n"
	for candidate, count := range counts {
		resultStr += fmt.Sprintf("Candidate %s: %d votes\n", candidate, count)
	}
	resultStr += fmt.Sprintf("Winner: Candidate %s with %d votes\n", winnerID, maxVotes)
	return resultStr
}
