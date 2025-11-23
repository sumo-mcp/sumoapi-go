package sumoapi

import (
	"encoding/json"
	"fmt"
)

// Match represents a sumo match.
type Match struct {
	// ID is optional because it's not returned by the APIs for listing rikishi matches.
	ID             *MatchID `json:"id,omitempty" jsonschema:"The unique identifier for the match in the format YYYYMM-D-NUM-EASTID-WESTID."`
	BashoID        BashoID  `json:"bashoId" jsonschema:"The ID of the basho (sumo tournament) in which the match took place, in the format YYYYMM."`
	Division       string   `json:"division" jsonschema:"The division in which the match took place."`
	Day            int      `json:"day" jsonschema:"The day of the basho (sumo tournament) on which the match took place."`
	MatchNumber    int      `json:"matchNo,omitempty" jsonschema:"The number of the match on the given day."`
	EastID         int      `json:"eastId,omitempty" jsonschema:"The unique identifier for the rikishi (sumo wrestler) on the east side."`
	EastShikona    string   `json:"eastShikona,omitempty" jsonschema:"The shikona (ring name) in English of the rikishi (sumo wrestler) on the east side."`
	EastRank       string   `json:"eastRank,omitempty" jsonschema:"The rank of the rikishi (sumo wrestler) on the east side."`
	WestID         int      `json:"westId,omitempty" jsonschema:"The unique identifier for the rikishi (sumo wrestler) on the west side."`
	WestShikona    string   `json:"westShikona,omitempty" jsonschema:"The shikona (ring name) in English of the rikishi (sumo wrestler) on the west side."`
	WestRank       string   `json:"westRank,omitempty" jsonschema:"The rank of the rikishi (sumo wrestler) on the west side."`
	WinnerID       int      `json:"winnerId,omitempty" jsonschema:"The unique identifier for the winning rikishi (sumo wrestler)."`
	WinnerEnglish  string   `json:"winnerEn,omitempty" jsonschema:"The shikona (ring name) in English of the winning rikishi (sumo wrestler)."`
	WinnerJapanese string   `json:"winnerJp,omitempty" jsonschema:"The shikona (ring name) in Japanese of the winning rikishi (sumo wrestler)."`
	Kimarite       string   `json:"kimarite,omitempty" jsonschema:"The kimarite (winning technique) used in the match."`
}

// MatchID represents the unique identifier for a sumo match.
type MatchID struct {
	BashoID
	Day         int
	MatchNumber int
	EastID      int
	WestID      int
}

func (m MatchID) String() string {
	return fmt.Sprintf("%s-%d-%d-%d-%d", m.BashoID.String(), m.Day, m.MatchNumber, m.EastID, m.WestID)
}

func (m MatchID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + m.String() + `"`), nil
}

func (m *MatchID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling MatchID: %w", err)
	}
	if s == "" { // Match ID is optional in some APIs.
		*m = MatchID{}
		return nil
	}
	var bashoID BashoID
	if len(s) < 11 {
		return fmt.Errorf("invalid MatchID format: %s", s)
	}
	if err := bashoID.UnmarshalJSON([]byte(`"` + s[0:6] + `"`)); err != nil {
		return fmt.Errorf("error parsing BashoID from MatchID: %w", err)
	}
	var day, matchNumber, eastID, westID int
	n, err := fmt.Sscanf(s[7:], "%d-%d-%d-%d", &day, &matchNumber, &eastID, &westID)
	if err != nil {
		return fmt.Errorf("error parsing day, match number, east ID, and west ID from MatchID: %w", err)
	}
	if n != 4 {
		return fmt.Errorf("invalid MatchID format: %s", s)
	}
	m.BashoID = bashoID
	m.Day = day
	m.MatchNumber = matchNumber
	m.EastID = eastID
	m.WestID = westID
	return nil
}
