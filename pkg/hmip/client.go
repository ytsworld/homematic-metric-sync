package hmip

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func CreateClient(authToken, accessPoint, userAgent, clientTokenSalt string) (*HmIPClient, error) {

	c := HmIPClient{
		AuthToken:       authToken,
		AccessPoint:     accessPoint,
		UserAgent:       userAgent,
		ClientTokenSalt: clientTokenSalt,
	}

	err := c.lookupHmIPHost()
	if err != nil {
		return nil, err
	}

	clientAuth := c.generateAuthClientToken()
	c.ClientAuthToken = clientAuth

	return &c, nil
}

func (c *HmIPClient) getClientCharacteristics() string {
	requestTemplate := `
	{
		"clientCharacteristics": {
				"apiVersion": "10",
				"applicationIdentifier": "homematicip-python",
				"applicationVersion": "1.0",
				"deviceManufacturer": "none",
				"deviceType": "Computer",
				"language": "en_US",
				"osType": "Linux",
				"osVersion": "5.15.0-46-generic"
		},
		"id": "%s"
	}
	`
	return fmt.Sprintf(requestTemplate, c.AccessPoint)
}

func (c *HmIPClient) generateAuthClientToken() string {
	sha := sha512.Sum512([]byte(fmt.Sprintf("%s%s", c.AccessPoint, c.ClientTokenSalt)))
	hexValue := hex.EncodeToString(sha[:])
	return strings.ToUpper(hexValue)
}

func (c *HmIPClient) lookupHmIPHost() error {
	client := http.Client{}

	reqBody := c.getClientCharacteristics()

	req, _ := http.NewRequest("POST", "https://lookup.homematic.com:48335/getHost", bytes.NewBuffer([]byte(reqBody)))
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return errors.New(fmt.Sprintf("Error doing lookup for homematic REST API. Code: %d Status: %s", resp.StatusCode, resp.Status))
	}

	var res HmIPLookupResponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return err
	}

	c.RESTEndpoint = res.UrlREST
	log.Printf("Received endpoint from lookup: %s", res.UrlREST)

	return nil
}
