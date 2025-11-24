package sumoapi

// Kimarite represents a sumo wrestling winning technique.
type Kimarite struct {
	Name      string     `json:"kimarite" jsonschema:"The kimarite (winning technique) name."`
	Count     int        `json:"count" jsonschema:"The number of times this kimarite (winning technique) has been used in a match."`
	LastUsage BashoDayID `json:"lastUsage" jsonschema:"The basho (sumo tournament) day (1-15) when this kimarite (winning technique) was last used in the format YYYYMM-{day}. A day value of 16 or higher represents an individual playoff match."`
}
