package sumoapi

// Measurement represents a rikishi's measurement in a specific basho.
type Measurement struct {
	ID        RikishiChangeID `json:"id" jsonschema:"The unique identifier for the measurement in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the rikishi in the API."`
	BashoID   BashoID         `json:"bashoId" jsonschema:"The ID of the basho (sumo tournament) in the format YYYYMM."`
	RikishiID int             `json:"rikishiId" jsonschema:"The unique identifier for the rikishi (sumo wrestler) in the API."`
	Height    float64         `json:"height,omitempty" jsonschema:"The height of the rikishi in centimeters measured on the beginning of the specified basho (sumo tournament)."`
	Weight    float64         `json:"weight,omitempty" jsonschema:"The weight of the rikishi in kilograms measured on the beginning of the specified basho (sumo tournament)."`
}
