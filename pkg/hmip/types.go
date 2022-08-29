package hmip

import "github.com/augurysys/timestamp"

type HmIPClient struct {
	AuthToken       string
	AccessPoint     string
	UserAgent       string
	RESTEndpoint    string
	ClientAuthToken string
	ClientTokenSalt string
	CurrentState    *HmIPCurrentStatus
	LastRequestTime timestamp.Timestamp
}

type HmIPLookupResponse struct {
	ApiVersion              string `json:"apiVersion"`
	PrimaryAccessPointId    string `json:"primaryAccessPointId"`
	RequestingAccessPointId string `json:"requestingAccessPointId"`
	UrlREST                 string `json:"urlREST"`
	UrlWebSocket            string `json:"urlWebSocket"`
}

type HmIPCurrentStatus struct {
	Groups  map[string]HmIPGroup  `json:"groups"`
	Devices map[string]HmIPDevice `json:"devices"`
}

type HmIPGroup struct {
	Id       string        `json:"id"`
	Label    string        `json:"label"`
	Type     string        `json:"type"`
	Channels []HmIPChannel `json:"channels"`
}

type HmIPChannel struct {
	DeviceId     string `json:"deviceId"`
	ChannelIndex int8   `json:"channelIndex"`
}

type HmIPDevice struct {
	Id                 string                           `json:"id"`
	Label              string                           `json:"label"`
	LastStatusUpdate   timestamp.Timestamp              `json:"lastStatusUpdate"`
	Type               string                           `json:"type"`
	FunctionalChannels map[string]HmIPFunctionalChannel `json:"functionalChannels"`
}

type HmIPFunctionalChannel struct {
	FunctionalChannelType   string  `json:"functionalChannelType"`
	ActualTemperature       float32 `json:"actualTemperature"`
	Humidity                int8    `json:"humidity"`
	VaporAmount             float32 `json:"vaporAmount"`
	EnergyCounter           float32 `json:"energyCounter"`
	CurrentPowerConsumption float32 `json:"currentPowerConsumption"`
	EnergyMeterMode         string  `json:"energyMeterMode"`
}
