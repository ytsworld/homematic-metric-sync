package hmip

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/augurysys/timestamp"
)

func (c *HmIPClient) FetchCurrentState() error {

	client := http.Client{}

	reqBody := c.getClientCharacteristics()

	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/hmip/home/getCurrentState", c.RESTEndpoint), bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("VERSION", "12")
	req.Header.Set("CLIENTAUTH", c.ClientAuthToken)
	req.Header.Set("AUTHTOKEN", c.AuthToken)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Error requesting data from HmIP; response code: %d, status: %s", resp.StatusCode, resp.Status))
	}

	var currentState HmIPCurrentStatus
	err = json.NewDecoder(resp.Body).Decode(&currentState)
	if err != nil {
		return err
	}

	/* 	dat, err := os.ReadFile("./local/getCurrentState.json")
	   	if err != nil {
	   		return err
	   	} */

	/* 	err = json.Unmarshal([]byte(dat), &currentState)
	   	if err != nil {
	   		return err
	   	} */

	c.CurrentState = &currentState
	c.LastRequestTime = *timestamp.Now()

	//For debugging
	stateString, err := json.Marshal(currentState)
	if err != nil {
		return err
	}

	f, err := os.OpenFile("./data/hmip-currentStateReponse.json",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(fmt.Sprintf("%s\n", stateString)); err != nil {
		log.Println(err)
	}

	return nil

}

func (c *HmIPClient) GetLastFetchedState() *HmIPCurrentStatus {
	return c.CurrentState
}
