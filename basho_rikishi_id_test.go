package sumoapi_test

import (
	"encoding/json"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/sumo-mcp/sumoapi-go"
)

func TestMarshalBashoRikishiIDToJSON(t *testing.T) {
	for _, tt := range []struct {
		name           string
		bashoRikishiID sumoapi.BashoRikishiID
		expectedJSON   string
	}{
		{
			name:           "double digit month",
			bashoRikishiID: sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1},
			expectedJSON:   `"202311-1"`,
		},
		{
			name:           "single digit month",
			bashoRikishiID: sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 1999, Month: 3}, RikishiID: 1},
			expectedJSON:   `"199903-1"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			b, err := json.Marshal(tt.bashoRikishiID)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(b)).To(Equal(tt.expectedJSON))
		})
	}

	t.Run("marshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.BashoRikishiID `json:"basho"`
		}
		w := wrapper{Basho: sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1}}
		b, err := json.Marshal(w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(b)).To(Equal(`{"basho":"202311-1"}`))
	})
}

func TestUnmarshalBashoRikishiIDFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name                   string
		jsonData               string
		expectedBashoRikishiID sumoapi.BashoRikishiID
		expectError            bool
	}{
		{
			name:                   "double digit month",
			jsonData:               `"202311-1"`,
			expectedBashoRikishiID: sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1},
			expectError:            false,
		},
		{
			name:                   "single digit month",
			jsonData:               `"199903-1"`,
			expectedBashoRikishiID: sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 1999, Month: 3}, RikishiID: 1},
			expectError:            false,
		},
		{
			name:                   "invalid format",
			jsonData:               `"20A911"`,
			expectedBashoRikishiID: sumoapi.BashoRikishiID{},
			expectError:            true,
		},
		{
			name:                   "wrong length",
			jsonData:               `"20231"`,
			expectedBashoRikishiID: sumoapi.BashoRikishiID{},
			expectError:            true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			var b sumoapi.BashoRikishiID
			err := json.Unmarshal([]byte(tt.jsonData), &b)
			if tt.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(b).To(Equal(tt.expectedBashoRikishiID))
			}
		})
	}

	t.Run("unmarshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			Basho sumoapi.BashoRikishiID `json:"basho"`
		}
		jsonData := `{"basho":"202311-1"}`
		var w wrapper
		err := json.Unmarshal([]byte(jsonData), &w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(w.Basho).To(Equal(sumoapi.BashoRikishiID{BashoID: sumoapi.BashoID{Year: 2023, Month: 11}, RikishiID: 1}))
	})
}
