package sumoapi

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	"github.com/google/jsonschema-go/jsonschema"
)

// RikishiChangeID represents the unique identifier for a rikishi change in a specific basho.
type RikishiChangeID struct {
	BashoID
	RikishiID int
}

func init() {
	typeSchemas[reflect.TypeFor[RikishiChangeID]()] = &jsonschema.Schema{Type: "string"}
}

func (r RikishiChangeID) String() string {
	return fmt.Sprintf("%s-%d", r.BashoID.String(), r.RikishiID)
}

func (r RikishiChangeID) MarshalJSON() ([]byte, error) {
	return []byte(`"` + r.String() + `"`), nil
}

func (r *RikishiChangeID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("error unmarshaling RikishiChangeID: %w", err)
	}
	if len(s) < 7 {
		return fmt.Errorf("invalid RikishiChangeID format: %s", s)
	}
	var bashoID BashoID
	if err := bashoID.UnmarshalJSON([]byte(`"` + s[0:6] + `"`)); err != nil {
		return fmt.Errorf("error parsing BashoID from RikishiChangeID: %w", err)
	}
	r.BashoID = bashoID
	rikishiID, err := strconv.Atoi(s[7:])
	if err != nil {
		return fmt.Errorf("error parsing RikishiID from RikishiChangeID: %w", err)
	}
	r.RikishiID = rikishiID
	return nil
}
