package sumoapi

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// BashoID represents the unique identifier for a basho.
type BashoID struct {
	Year  int
	Month int // Month is 1-12.
}

// BashoDayID represents a specific day within a basho (1-15), or a playoff match starting from 16.
type BashoDayID struct {
	BashoID
	Day int
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

func (b BashoDayID) String() string {
	return fmt.Sprintf("%s-%d", b.BashoID.String(), b.Day)
}

func (b BashoDayID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + b.String() + `"`), nil
}

func (b *BashoDayID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling BashoDayID: %w", err)
	}
	if len(s) < 8 {
		return fmt.Errorf("invalid BashoDayID format: %s", s)
	}
	var bashoID BashoID
	if err := bashoID.UnmarshalJSON([]byte(`"` + s[0:6] + `"`)); err != nil {
		return fmt.Errorf("error parsing BashoID from BashoDayID: %w", err)
	}
	day, err := strconv.Atoi(s[7:])
	if err != nil {
		return fmt.Errorf("error parsing day from BashoDayID: %w", err)
	}
	b.BashoID = bashoID
	b.Day = day
	return nil
}
