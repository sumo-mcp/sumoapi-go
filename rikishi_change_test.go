package sumoapi_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestMarshalRikishiChangeIDToJSON(t *testing.T) {
	for _, tt := range []struct {
		name            string
		rikishiChangeID sumoapi.RikishiChangeID
		expectedJSON    string
	}{
		{
			name:            "double digit month",
			rikishiChangeID: sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1},
			expectedJSON:    `"202311-1"`,
		},
		{
			name:            "single digit month",
			rikishiChangeID: sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 1999, Month: 3}, RikishiID: 1},
			expectedJSON:    `"199903-1"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			b, err := json.Marshal(tt.rikishiChangeID)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(b)).To(Equal(tt.expectedJSON))
		})
	}

	t.Run("marshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.RikishiChangeID `json:"basho"`
		}
		w := wrapper{Basho: sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1}}
		b, err := json.Marshal(w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(b)).To(Equal(`{"basho":"202311-1"}`))
	})
}

func TestUnmarshalRikishiChangeIDFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name                    string
		jsonData                string
		expectedRikishiChangeID sumoapi.RikishiChangeID
		expectError             bool
	}{
		{
			name:                    "double digit month",
			jsonData:                `"202311-1"`,
			expectedRikishiChangeID: sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1},
			expectError:             false,
		},
		{
			name:                    "single digit month",
			jsonData:                `"199903-1"`,
			expectedRikishiChangeID: sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 1999, Month: 3}, RikishiID: 1},
			expectError:             false,
		},
		{
			name:                    "invalid format",
			jsonData:                `"20A911"`,
			expectedRikishiChangeID: sumoapi.RikishiChangeID{},
			expectError:             true,
		},
		{
			name:                    "wrong length",
			jsonData:                `"20231"`,
			expectedRikishiChangeID: sumoapi.RikishiChangeID{},
			expectError:             true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			var b sumoapi.RikishiChangeID
			err := json.Unmarshal([]byte(tt.jsonData), &b)
			if tt.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(b).To(Equal(tt.expectedRikishiChangeID))
			}
		})
	}

	t.Run("unmarshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.RikishiChangeID `json:"basho"`
		}
		jsonData := `{"basho":"202311-1"}`
		var w wrapper
		err := json.Unmarshal([]byte(jsonData), &w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(w.Basho).To(Equal(sumoapi.RikishiChangeID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1}))
	})
}
