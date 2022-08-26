package hmip

import "github.com/augurysys/timestamp"

type HmIPClient struct {
	AuthToken    string
	AccessPoint  string
	UserAgent    string
	CurrentState *HmIPCurrentStatus
}

type HmIPCurrentStatus struct {
	Groups  map[string]HmIPGroup  `json:"groups"`
	Devices map[string]HmIPDevice `json:"devices"`
}

type HmIPGroup struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type HmIPDevice struct {
	Id                 string                           `json:"id"`
	Label              string                           `json:"label"`
	LastStatusUpdate   timestamp.Timestamp              `json:"lastStatusUpdate"`
	Type               string                           `json:"type"`
	FunctionalChannels map[string]HmIPFunctionalChannel `json:"functionalChannels"`
}

type HmIPFunctionalChannel struct {
	FunctionalChannelType string  `json:"functionalChannelType"`
	ActualTemperature     float32 `json:"actualTemperature"`
	Humidity              int8    `json:"humidity"`
	VaporAmount           float32 `json:"vaporAmount"`
}
