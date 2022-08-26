package main

import (
	"fmt"

	"github.com/ytsworld/homematic-metric-sync/pkg/hmip"
)

func init() {

}

var (
	deviceTypesForHistorization []string = []string{"WALL_MOUNTED_THERMOSTAT_PRO", "PLUGABLE_SWITCH_MEASURING", "TEMPERATURE_HUMIDITY_SENSOR_OUTDOOR"}
)

func main() {

	fmt.Printf("Hello.\n")

	c, err := hmip.CreateClient()
	if err != nil {
		panic(err)
	}

	err = c.FetchCurrentState()
	if err != nil {
		panic(err)
	}

	state := c.GetLastFetchedState()

	for k, device := range state.Devices {

		if contains(deviceTypesForHistorization, device.Type) {
			room := c.SearchRoomForDevice(k)
			fmt.Printf("Device ID: %s, type: %s, label: %s, room: %s\n", k, device.Type, device.Label, room)
			if device.Type == "WALL_MOUNTED_THERMOSTAT_PRO" {
				for _, fc := range device.FunctionalChannels {
					if fc.FunctionalChannelType == "WALL_MOUNTED_THERMOSTAT_PRO_CHANNEL" {
						fmt.Printf(" - temp: %f, humidty: %d\n", fc.ActualTemperature, fc.Humidity)
					}
				}
			}
			if device.Type == "PLUGABLE_SWITCH_MEASURING" {
				for _, fc := range device.FunctionalChannels {
					if fc.FunctionalChannelType == "SWITCH_MEASURING_CHANNEL" {
						fmt.Printf(" - counter: %f, consumption: %f, mode: %s\n", fc.EnergyCounter, fc.CurrentPowerConsumption, fc.EnergyMeterMode)
					}
				}
			}
			if device.Type == "TEMPERATURE_HUMIDITY_SENSOR_OUTDOOR" {
				for _, fc := range device.FunctionalChannels {
					if fc.FunctionalChannelType == "CLIMATE_SENSOR_CHANNEL" {
						fmt.Printf(" - temp: %f, humidty: %d\n", fc.ActualTemperature, fc.Humidity)
					}
				}
			}
		}

	}

	/* 	str, err := json.Marshal(currentState)
	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Println("output: ", string(str)) */

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
