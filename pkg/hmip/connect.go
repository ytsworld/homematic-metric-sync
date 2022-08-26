package hmip

import (
	"encoding/json"
	"os"
)

func (c *HmIPClient) FetchCurrentState() error {

	currentState := HmIPCurrentStatus{}

	dat, err := os.ReadFile("./local/getCurrentState.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(dat), &currentState)
	if err != nil {
		return err
	}

	c.CurrentState = &currentState

	return nil

}

func (c *HmIPClient) GetLastFetchedState() *HmIPCurrentStatus {
	return c.CurrentState
}
