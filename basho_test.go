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

func TestMarshalBashoDayIDToJSON(t *testing.T) {
	for _, tt := range []struct {
		name         string
		bashoDayID   sumoapi.BashoDayID
		expectedJSON string
	}{
		{
			name: "double digit month, single digit day",
			bashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 2023, Month: 11},
				Day:     1,
			},
			expectedJSON: `"202311-1"`,
		},
		{
			name: "single digit month, double digit day",
			bashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 1999, Month: 3},
				Day:     15,
			},
			expectedJSON: `"199903-15"`,
		},
		{
			name: "playoff day",
			bashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
				Day:     16,
			},
			expectedJSON: `"202501-16"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			b, err := json.Marshal(tt.bashoDayID)
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(string(b)).To(Equal(tt.expectedJSON))
		})
	}

	t.Run("marshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			BashoDay sumoapi.BashoDayID `json:"bashoDay"`
		}
		w := wrapper{
			BashoDay: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 2023, Month: 11},
				Day:     1,
			},
		}
		b, err := json.Marshal(w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(string(b)).To(Equal(`{"bashoDay":"202311-1"}`))
	})
}

func TestUnmarshalBashoDayIDFromJSON(t *testing.T) {
	for _, tt := range []struct {
		name               string
		jsonData           string
		expectedBashoDayID sumoapi.BashoDayID
		expectError        bool
	}{
		{
			name:     "double digit month, single digit day",
			jsonData: `"202311-1"`,
			expectedBashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 2023, Month: 11},
				Day:     1,
			},
			expectError: false,
		},
		{
			name:     "single digit month, double digit day",
			jsonData: `"199903-15"`,
			expectedBashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 1999, Month: 3},
				Day:     15,
			},
			expectError: false,
		},
		{
			name:     "playoff day",
			jsonData: `"202501-16"`,
			expectedBashoDayID: sumoapi.BashoDayID{
				BashoID: sumoapi.BashoID{Year: 2025, Month: 1},
				Day:     16,
			},
			expectError: false,
		},
		{
			name:               "invalid basho format",
			jsonData:           `"20A311-1"`,
			expectedBashoDayID: sumoapi.BashoDayID{},
			expectError:        true,
		},
		{
			name:               "wrong length - too short",
			jsonData:           `"202311"`,
			expectedBashoDayID: sumoapi.BashoDayID{},
			expectError:        true,
		},
		{
			name:               "missing day",
			jsonData:           `"202311-"`,
			expectedBashoDayID: sumoapi.BashoDayID{},
			expectError:        true,
		},
		{
			name:               "invalid day - not a number",
			jsonData:           `"202311-X"`,
			expectedBashoDayID: sumoapi.BashoDayID{},
			expectError:        true,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			g := NewWithT(t)
			var b sumoapi.BashoDayID
			err := json.Unmarshal([]byte(tt.jsonData), &b)
			if tt.expectError {
				g.Expect(err).To(HaveOccurred())
			} else {
				g.Expect(err).ToNot(HaveOccurred())
				g.Expect(b).To(Equal(tt.expectedBashoDayID))
			}
		})
	}

	t.Run("unmarshal inside struct", func(t *testing.T) {
		g := NewWithT(t)
		type wrapper struct {
			BashoDay sumoapi.BashoDayID `json:"bashoDay"`
		}
		jsonData := `{"bashoDay":"202311-1"}`
		var w wrapper
		err := json.Unmarshal([]byte(jsonData), &w)
		g.Expect(err).ToNot(HaveOccurred())
		g.Expect(w.BashoDay).To(Equal(sumoapi.BashoDayID{
			BashoID: sumoapi.BashoID{Year: 2023, Month: 11},
			Day:     1,
		}))
	})
}
