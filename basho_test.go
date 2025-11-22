package sumoapi_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestMarshalBashoIDToJSON(t *testing.T) {
	for _, tt := range []struct {
		name         string
		bashoID      sumoapi.BashoID
		expectedJSON string
	}{
		{
			name:         "double digit month",
			bashoID:      sumoapi.BashoID{Year: 2023, Month: 11},
			expectedJSON: `"202311"`,
		},
		{
			name:         "single digit month",
			bashoID:      sumoapi.BashoID{Year: 1999, Month: 3},
			expectedJSON: `"199903"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			b, err := json.Marshal(tt.bashoID)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(b)).To(Equal(tt.expectedJSON))
		})
	}

	t.Run("marshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.BashoID `json:"basho"`
		}
		w := wrapper{Basho: sumoapi.BashoID{Year: 2023, Month: 11}}
		b, err := json.Marshal(w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(b)).To(Equal(`{"basho":"202311"}`))
	})
}

func TestUnmarshalBashoIDFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name            string
		jsonData        string
		expectedBashoID sumoapi.BashoID
		expectError     bool
	}{
		{
			name:            "double digit month",
			jsonData:        `"202311"`,
			expectedBashoID: sumoapi.BashoID{Year: 2023, Month: 11},
			expectError:     false,
		},
		{
			name:            "single digit month",
			jsonData:        `"199903"`,
			expectedBashoID: sumoapi.BashoID{Year: 1999, Month: 3},
			expectError:     false,
		},
		{
			name:            "invalid format",
			jsonData:        `"20A911"`,
			expectedBashoID: sumoapi.BashoID{},
			expectError:     true,
		},
		{
			name:            "wrong length",
			jsonData:        `"20231"`,
			expectedBashoID: sumoapi.BashoID{},
			expectError:     true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			var b sumoapi.BashoID
			err := json.Unmarshal([]byte(tt.jsonData), &b)
			if tt.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(b).To(Equal(tt.expectedBashoID))
			}
		})
	}

	t.Run("unmarshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.BashoID `json:"basho"`
		}
		jsonData := `{"basho":"202311"}`
		var w wrapper
		err := json.Unmarshal([]byte(jsonData), &w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(w.Basho).To(Equal(sumoapi.BashoID{Year: 2023, Month: 11}))
	})
}
