package systems

// VotingSystem interface
type VotingSystem interface {
	RegisterVoter(voter Voter) error
	CastVote(voterID int, vote interface{}) error
	CountVotes() (map[string]int, error)
	ReportResults() string
}

// Example of a specific voting mechanism
type FirstPastThePost struct {
	Voters []Voter
	Votes  map[int]string // voterID to candidate
}

// Implementation of VotingSystem for FirstPastThePost
func (f *FirstPastThePost) RegisterVoter(voter Voter) error {
	// Implementation specific to FPTP
}

func (f *FirstPastThePost) CastVote(voterID int, vote interface{}) error {
	// Implementation specific to FPTP
}

func (f *FirstPastThePost) CountVotes() (map[string]int, error) {
	// Implementation specific to FPTP
}

func (f *FirstPastThePost) ReportResults() string {
	// Implementation specific to FPTP
}
