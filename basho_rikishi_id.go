package sumoapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BashoRikishiID represents the unique identifier for a Rikishi in a specific Basho (sumo tournament).
type BashoRikishiID struct {
	BashoID
	RikishiID int
}

func (b BashoRikishiID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + fmt.Sprintf("%04d%02d-%d", b.Year, b.Month, b.RikishiID) + `"`), nil
}

func (b *BashoRikishiID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling BashoRikishiID: %w", err)
	}
	if len(s) < 7 {
		return fmt.Errorf("invalid BashoRikishiID format: %s", s)
	}
	year, err := strconv.Atoi(s[0:4])
	if err != nil {
		return fmt.Errorf("error parsing year from BashoRikishiID: %w", err)
	}
	month, err := strconv.Atoi(s[4:6])
	if err != nil {
		return fmt.Errorf("error parsing month from BashoRikishiID: %w", err)
	}
	b.Year = year
	b.Month = month
	rikishiID, err := strconv.Atoi(s[7:])
	if err != nil {
		return fmt.Errorf("error parsing RikishiID from BashoRikishiID: %w", err)
	}
	b.RikishiID = rikishiID
	return nil
}
