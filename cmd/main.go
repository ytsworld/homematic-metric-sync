package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ytsworld/homematic-metric-sync/pkg/hmip"
	"github.com/ytsworld/homematic-metric-sync/pkg/sync"
)

func init() {

}

func main() {

	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		configFile = "./hmip_sync.yaml"
	}

	config, err := ReadConfig(configFile)
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	c, err := hmip.CreateClient(config.HmIP.AuthToken, config.HmIP.AccessPoint, config.HmIP.UserAgent)
	if err != nil {
		log.Fatalf("Error connecting to HmIP: %s", err)
	}

	influxClient := sync.CreateClient(config.Influx.Url, config.Influx.Token, config.Influx.Organization, config.Influx.Bucket)

	for {
		currentDate := time.Now().Format("2006-01-02")

		err = c.FetchCurrentState()
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
			continue
		}

		metrics := sync.ConvertHmIPStateToMetrics(c)

		metricsJson, err := json.Marshal(metrics)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
			continue
		}

		f, err := os.OpenFile(fmt.Sprintf("./data/%s-metrics.log", currentDate),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
			time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
			continue
		}
		defer f.Close()

		if _, err := f.WriteString(fmt.Sprintf("%s\n", metricsJson)); err != nil {
			log.Println(err)
			time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
			continue
		}

		err = influxClient.WriteMetricsToInflux(metrics)
		if err != nil {
			log.Println(fmt.Sprintf("Error writing metrics to influx: %s.", err))
			time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
			continue
		}

		log.Printf("Persisted data from %d devices", len(metrics.Metrics))
		time.Sleep(time.Duration(config.HmIP.PollInterval) * time.Second)
	}

}
