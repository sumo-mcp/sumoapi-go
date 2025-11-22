package sumoapi

// Rank represents a sumo wrestler's rank in a specific Basho (sumo tournament).
type Rank struct {
	ID                BashoRikishiID `json:"id" jsonschema:"The unique identifier for the Rank in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the Rikishi in the API."`
	BashoID           BashoID        `json:"bashoId" jsonschema:"The ID of the Basho (sumo tournament) in the format YYYYMM."`
	RikishiID         int            `json:"rikishiId" jsonschema:"The unique identifier for the Rikishi (sumo wrestler) in the API."`
	HumanReadableName string         `json:"rank,omitempty" jsonschema:"The human-readable name of the rank (e.g., Maegashira 1 East, Ozeki 2 West) of the Rikishi in the specific Basho (sumo tournament)."`
	NumericName       int            `json:"rankValue,omitempty" jsonschema:"The numeric name of the rank of the Rikishi in the specific Basho (sumo tournament)."`
}
