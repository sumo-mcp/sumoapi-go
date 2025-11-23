package sumoapi_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestMarshalMatchIDToJSON(t *testing.T) {
	for _, tt := range []struct {
		name         string
		matchID      sumoapi.MatchID
		expectedJSON string
	}{
		{
			name: "double digit month",
			matchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 2023, Month: 11},
				Day:         1,
				MatchNumber: 5,
				EastID:      45,
				WestID:      123,
			},
			expectedJSON: `"202311-1-5-45-123"`,
		},
		{
			name: "single digit month",
			matchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 1999, Month: 3},
				Day:         15,
				MatchNumber: 42,
				EastID:      100,
				WestID:      200,
			},
			expectedJSON: `"199903-15-42-100-200"`,
		},
		{
			name: "large match number",
			matchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 2025, Month: 1},
				Day:         10,
				MatchNumber: 999,
				EastID:      1,
				WestID:      2,
			},
			expectedJSON: `"202501-10-999-1-2"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			b, err := json.Marshal(tt.matchID)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(b)).To(Equal(tt.expectedJSON))
		})
	}

	t.Run("marshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Match sumoapi.MatchID `json:"match"`
		}
		w := wrapper{
			Match: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 2023, Month: 11},
				Day:         1,
				MatchNumber: 5,
				EastID:      45,
				WestID:      123,
			},
		}
		b, err := json.Marshal(w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(b)).To(Equal(`{"match":"202311-1-5-45-123"}`))
	})
}

func TestUnmarshalMatchIDFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name            string
		jsonData        string
		expectedMatchID sumoapi.MatchID
		expectError     bool
	}{
		{
			name:     "double digit month",
			jsonData: `"202311-1-5-45-123"`,
			expectedMatchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 2023, Month: 11},
				Day:         1,
				MatchNumber: 5,
				EastID:      45,
				WestID:      123,
			},
			expectError: false,
		},
		{
			name:     "single digit month",
			jsonData: `"199903-15-42-100-200"`,
			expectedMatchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 1999, Month: 3},
				Day:         15,
				MatchNumber: 42,
				EastID:      100,
				WestID:      200,
			},
			expectError: false,
		},
		{
			name:     "large match number",
			jsonData: `"202501-10-999-1-2"`,
			expectedMatchID: sumoapi.MatchID{
				BashoID:     sumoapi.BashoID{Year: 2025, Month: 1},
				Day:         10,
				MatchNumber: 999,
				EastID:      1,
				WestID:      2,
			},
			expectError: false,
		},
		{
			name:            "empty string",
			jsonData:        `""`,
			expectedMatchID: sumoapi.MatchID{},
		},
		{
			name:            "invalid basho format",
			jsonData:        `"20A311-1-5-45-123"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "wrong length - too short",
			jsonData:        `"202311-1"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "missing day",
			jsonData:        `"202311--5-45-123"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "missing match number",
			jsonData:        `"202311-1--45-123"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "missing east ID",
			jsonData:        `"202311-1-5--123"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "missing west ID",
			jsonData:        `"202311-1-5-45-"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
		{
			name:            "invalid day - not a number",
			jsonData:        `"202311-X-5-45-123"`,
			expectedMatchID: sumoapi.MatchID{},
			expectError:     true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			var m sumoapi.MatchID
			err := json.Unmarshal([]byte(tt.jsonData), &m)
			if tt.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(m).To(Equal(tt.expectedMatchID))
			}
		})
	}

	t.Run("unmarshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Match sumoapi.MatchID `json:"match"`
		}
		jsonData := `{"match":"202311-1-5-45-123"}`
		var w wrapper
		err := json.Unmarshal([]byte(jsonData), &w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(w.Match).To(Equal(sumoapi.MatchID{
			BashoID:     sumoapi.BashoID{Year: 2023, Month: 11},
			Day:         1,
			MatchNumber: 5,
			EastID:      45,
			WestID:      123,
		}))
	})
}
