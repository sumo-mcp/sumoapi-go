package sumoapi

// Measurement represents a sumo wrestler's measurement in a specific Basho (sumo tournament).
type Measurement struct {
	ID        RikishiChangeID `json:"id" jsonschema:"The unique identifier for the Measurement in the format {bashoID}-{rikishiID} where {bashoID} is in the format YYYYMM and {rikishiID} is the unique identifier for the Rikishi in the API."`
	BashoID   BashoID         `json:"bashoId" jsonschema:"The ID of the Basho (sumo tournament) in the format YYYYMM."`
	RikishiID int             `json:"rikishiId" jsonschema:"The unique identifier for the Rikishi (sumo wrestler) in the API."`
	Height    float64         `json:"height,omitempty" jsonschema:"The height of the Rikishi in centimeters measured on the beginning of the specified Basho (sumo tournament)."`
	Weight    float64         `json:"weight,omitempty" jsonschema:"The weight of the Rikishi in kilograms measured on the beginning of the specified Basho (sumo tournament)."`
}
