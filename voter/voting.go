package voting

import (
	"math/rand"
	"time"
)

// Voter struct to represent a voter
type Voter struct {
	ID                 int
	Age                int
	IncomeLevel        string
	Education          string
	PoliticalAlignment string
	Preferences        map[string]float64
}

type Candidate struct {
	ID         int
	Party      string
	Age        int
	Experience string
	Policy     map[string]float64
}

type Vote struct {
	VoterID     int
	CandidateID []float32
}

// VotingSystem interface
type VotingSystem interface {
	RegisterVoter(voter Voter) error
	CastVote(voter Voter, candidate []Candidate) (Vote, error)
	CountVotes(votes []Vote) (map[string]int, error)
	ReportResults() string
}

func GenerateVoter(voterID int) Voter {
	incomeLevels := []string{"low", "middle", "high"}
	educationLevels := []string{"high school", "college", "graduate"}
	politicalAlignments := []string{"left", "center", "right"}
	socialIssues := []string{"environment", "economy", "immigration"}

	// Random voter profile
	voter := Voter{
		ID:                 voterID,       // Random ID between 0 and 999
		Age:                rand.Intn(80), // Random age between 0 and 79
		IncomeLevel:        incomeLevels[rand.Intn(len(incomeLevels))],
		Education:          educationLevels[rand.Intn(len(educationLevels))],
		PoliticalAlignment: politicalAlignments[rand.Intn(len(politicalAlignments))],
		Preferences:        GenerateRandomPreferences(socialIssues),
	}
	return voter
}

// Function to generate random preferences for social issues
func GenerateRandomPreferences(issues []string) map[string]float64 {
	preferences := make(map[string]float64)
	for _, issue := range issues {
		preferences[issue] = rand.Float64() // Random preference between 0.0 - 1.0
	}
	return preferences
}

// Function to generate a slice of random voters
func GenerateRandomVoters(count int) []Voter {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	voters := make([]Voter, count)
	for i := 0; i < count; i++ {
		voters[i] = GenerateVoter(i)
	}
	return voters
}

func GenerateCandidate(candidateID int) Candidate {
	politicalAlignments := []string{"left", "center", "right"}
	socialIssues := []string{"environment", "economy", "immigration"}

	candidate := Candidate{
		ID:         candidateID, // Random ID between 0 and 999
		Party:      politicalAlignments[rand.Intn(len(politicalAlignments))],
		Age:        rand.Intn(80), // Random age between 0 and 79
		Experience: "Senator",
		Policy:     generateRandomPolicy(socialIssues),
	}
	return candidate
}

func generateRandomPolicy(issues []string) map[string]float64 {
	policy := make(map[string]float64)
	for _, issue := range issues {
		policy[issue] = rand.Float64() // Random policy score between 0.0 - 1.0
	}
	return policy
}

func GenerateRandomCandidates(count int) []Candidate {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	candidates := make([]Candidate, count)
	for i := 0; i < count; i++ {
		candidates[i] = GenerateCandidate(i)
	}
	return candidates
}
