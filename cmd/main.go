package main

import (
	"fmt"

	"github.com/ytsworld/homematic-metric-sync/pkg/hmip"
)

func init() {

}

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

	for k, v := range state.Devices {
		fmt.Printf("Device ID: %s, label: %s\n", k, v.Label)
	}

	/* 	str, err := json.Marshal(currentState)
	   	if err != nil {
	   		panic(err)
	   	}

	   	fmt.Println("output: ", string(str)) */

}
