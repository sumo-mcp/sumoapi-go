package sumoapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"time"

	"github.com/google/jsonschema-go/jsonschema"
)

// Basho represents a sumo tournament.
type Basho struct {
	ID            BashoID      `json:"date" jsonschema:"The unique identifier for the basho (sumo tournament), in the format YYYYMM."`
	StartDate     *time.Time   `json:"startDate,omitempty" jsonschema:"The starting date of the basho (sumo tournament)."`
	EndDate       *time.Time   `json:"endDate,omitempty" jsonschema:"The ending date of the basho (sumo tournament)."`
	Yusho         []BashoPrize `json:"yusho,omitempty" jsonschema:"A list of yusho (tournament championship) prizes awarded to rikishi (sumo wrestlers) in the basho (sumo tournament)."`
	SpecialPrizes []BashoPrize `json:"specialPrizes,omitempty" jsonschema:"A list of special prizes awarded to rikishi (sumo wrestlers) in the basho (sumo tournament)."`
	Torikumi      []Match      `json:"torikumi,omitempty" jsonschema:"A torikumi (bout schedule) that took or will take place for a specific day of a specific division of the basho (sumo tournament)."`
}

// BashoPrize represents a prize awarded to a rikishi in a basho.
// It can be a yusho or a special prize.
type BashoPrize struct {
	Type            string `json:"type" jsonschema:"The type of prize. When the prize is a yusho (tournament championship), the value is the name of the tournament division. When the prize is a special prize, the value is one of 'Shukun-sho' (outstanding performance), 'Kanto-sho' (fighting spirit), or 'Gino-sho' (technique)."`
	RikishiID       int    `json:"rikishiId" jsonschema:"The unique identifier of the rikishi (sumo wrestler) who received the prize."`
	ShikonaEnglish  string `json:"shikonaEn,omitempty" jsonschema:"The shikona (ring name) of the rikishi (sumo wrestler) in English."`
	ShikonaJapanese string `json:"shikonaJp,omitempty" jsonschema:"The shikona (ring name) of the rikishi (sumo wrestler) in Japanese."`
}

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

func init() {
	typeSchemas[reflect.TypeFor[BashoID]()] = &jsonschema.Schema{Type: "string"}
	typeSchemas[reflect.TypeFor[BashoDayID]()] = &jsonschema.Schema{Type: "string"}
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
