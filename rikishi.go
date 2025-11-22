package sumoapi

import "time"

// Rikishi represents a sumo wrestler.
type Rikishi struct {
	ID                 int           `json:"id" jsonschema:"The unique identifier for the Rikishi in the API."`
	SumoDBID           int           `json:"sumodbId,omitempty" jsonschema:"The SumoDB ID of the Rikishi."`
	OfficialID         int           `json:"nskId,omitempty" jsonschema:"The official Nihon Sumo Kyokai (Japan Sumo Association) ID of the Rikishi."`
	ShikonaEnglish     string        `json:"shikonaEn,omitempty" jsonschema:"The shikona (ring name) of the Rikishi in English."`
	ShikonaJapanese    string        `json:"shikonaJp,omitempty" jsonschema:"The shikona (ring name) of the Rikishi in Japanese."`
	CurrentRank        string        `json:"currentRank,omitempty" jsonschema:"The current rank of the Rikishi."`
	Heya               string        `json:"heya,omitempty" jsonschema:"The heya (stable) of the Rikishi."`
	BirthDate          *time.Time    `json:"birthDate,omitempty" jsonschema:"The birth date of the Rikishi."`
	Shusshin           string        `json:"shusshin,omitempty" jsonschema:"The place of birth of the Rikishi."`
	Height             float64       `json:"height,omitempty" jsonschema:"The height of the Rikishi in centimeters."`
	Weight             float64       `json:"weight,omitempty" jsonschema:"The weight of the Rikishi in kilograms."`
	Debut              *BashoID      `json:"debut,omitempty" jsonschema:"The ID of the Basho when the Rikishi made their debut in the format YYYYMM."`
	Intai              *time.Time    `json:"intai,omitempty" jsonschema:"The retirement date of the Rikishi, if retired."`
	RankHistory        []Rank        `json:"rankHistory,omitempty" jsonschema:"The historical rank records of the Rikishi over time. Each record is the rank of the Rikishi in a specific Basho (sumo tournament)."`
	ShikonaHistory     []Shikona     `json:"shikonaHistory,omitempty" jsonschema:"The historical shikona (ring name) records of the Rikishi over time. Each record is the shikona of the Rikishi in a specific Basho (sumo tournament)."`
	MeasurementHistory []Measurement `json:"measurementHistory,omitempty" jsonschema:"The historical measurement records of the Rikishi over time. Each record includes height and weight of the Rikishi in a specific Basho (sumo tournament)."`
	CreatedAt          *time.Time    `json:"createdAt,omitempty" jsonschema:"The timestamp when the Rikishi record was created in the API."`
	UpdatedAt          *time.Time    `json:"updatedAt,omitempty" jsonschema:"The timestamp when the Rikishi record was last updated in the API."`
}
