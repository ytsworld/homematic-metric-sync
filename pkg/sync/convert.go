package sync

import (
	"github.com/ytsworld/homematic-metric-sync/pkg/hmip"
)

var (
	deviceTypesForHistorization = []string{"WALL_MOUNTED_THERMOSTAT_PRO", "PLUGABLE_SWITCH_MEASURING", "TEMPERATURE_HUMIDITY_SENSOR_OUTDOOR"}
)

func ConvertHmIPStateToMetrics(client *hmip.HmIPClient) *Metrics {

	state := client.GetLastFetchedState()
	metrics := Metrics{}
	metrics.RequestTime = client.LastRequestTime

	for id, device := range state.Devices {

		if Contains(deviceTypesForHistorization, device.Type) {
			room := client.SearchRoomForDevice(id)

			metric := Metric{}
			metric.DeviceId = id
			metric.DeviceLabel = device.Label
			metric.DeviceType = device.Type
			metric.Location = room
			metric.LastStatusUpdate = device.LastStatusUpdate

			if device.Type == "WALL_MOUNTED_THERMOSTAT_PRO" || device.Type == "TEMPERATURE_HUMIDITY_SENSOR_OUTDOOR" {
				metric.MetricType = "Climate"
				climate := ClimateMetric{}
				for _, fc := range device.FunctionalChannels {
					if fc.FunctionalChannelType == "WALL_MOUNTED_THERMOSTAT_PRO_CHANNEL" || fc.FunctionalChannelType == "CLIMATE_SENSOR_CHANNEL" {
						climate.Humidity = fc.Humidity
						climate.Temperature = fc.ActualTemperature
					}
				}
				metric.Climate = &climate
			}
			if device.Type == "PLUGABLE_SWITCH_MEASURING" {
				metric.MetricType = "Energy"
				em := EnergyMetric{}
				for _, fc := range device.FunctionalChannels {
					if fc.FunctionalChannelType == "SWITCH_MEASURING_CHANNEL" {
						em.TotalCounter = fc.EnergyCounter
						em.MeasuringMode = fc.EnergyMeterMode
						em.Current = fc.CurrentPowerConsumption
					}
				}
				metric.Energy = &em
			}

			metrics.Metrics = append(metrics.Metrics, &metric)
		}
	}

	return &metrics
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
