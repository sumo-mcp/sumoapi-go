package sumoapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BashoID represents the unique identifier for a Basho (sumo tournament).
type BashoID struct {
	Year  int
	Month int // Month is 1-12.
}

func (b BashoID) String() string {
	return fmt.Sprintf("%04d%02d", b.Year, b.Month)
}

func (b BashoID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + b.String() + `"`), nil
}

func (b *BashoID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling BashoID: %w", err)
	}
	if len(s) != 6 {
		return fmt.Errorf("invalid BashoID format: %s", s)
	}
	year, err := strconv.Atoi(s[0:4])
	if err != nil {
		return fmt.Errorf("error parsing year from BashoID: %w", err)
	}
	month, err := strconv.Atoi(s[4:6])
	if err != nil {
		return fmt.Errorf("error parsing month from BashoID: %w", err)
	}
	b.Year = year
	b.Month = month
	return nil
}
