package sync

import "github.com/augurysys/timestamp"

type Metrics struct {
	RequestTime timestamp.Timestamp `json:"requestTime"`
	Metrics     []*Metric           `json:"metrics"`
}

type Metric struct {
	DeviceId         string              `json:"deviceId"`
	DeviceLabel      string              `json:"deviceLabel"`
	DeviceType       string              `json:"deviceType"`
	MetricType       string              `json:"metricType"`
	Location         string              `json:"location"`
	LastStatusUpdate timestamp.Timestamp `json:"lastStatusUpdate"`
	Climate          *ClimateMetric      `json:"climate,omitempty"`
	Energy           *EnergyMetric       `json:"energy,omitempty"`
}

type ClimateMetric struct {
	Temperature float32 `json:"temperature"`
	Humidity    int8    `json:"humidity"`
}

type EnergyMetric struct {
	TotalCounter  float32 `json:"totalCounter"`
	MeasuringMode string  `json:"measuringMode"`
	Current       float32 `json:"current"`
}
