package sync

import (
	"context"
	"fmt"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

type InfluxClient struct {
	Url          string
	Token        string
	Organization string
	Bucket       string
}

func CreateClient(url string, token string, organization string, bucket string) *InfluxClient {
	return &InfluxClient{
		Url:          url,
		Token:        token,
		Organization: organization,
		Bucket:       bucket,
	}
}

func (c *InfluxClient) WriteMetricsToInflux(metrics *Metrics) error {
	client := influxdb2.NewClient(c.Url, c.Token)
	defer client.Close()

	writeAPI := client.WriteAPIBlocking(c.Organization, c.Bucket)

	for _, m := range metrics.Metrics {
		err := writeMetric(writeAPI, m)
		if err != nil {
			return err
		}
	}

	return nil
}

func writeMetric(writeAPI api.WriteAPIBlocking, metric *Metric) error {
	measurement := fmt.Sprintf("%s: %s", metric.Location, metric.DeviceLabel)

	tags := map[string]string{
		"location":   metric.Location,
		"deviceId":   metric.DeviceId,
		"metricType": metric.MetricType,
		"label":      metric.DeviceLabel,
	}

	var fields map[string]interface{}

	if metric.MetricType == "Energy" {
		tags["mode"] = metric.Energy.MeasuringMode
		fields = map[string]interface{}{"total": metric.Energy.TotalCounter, "current": metric.Energy.Current}
	} else if metric.MetricType == "Climate" {
		fields = map[string]interface{}{"humidity": metric.Climate.Humidity, "temperature": metric.Climate.Temperature} //TODO
	} else {
		log.Printf("ERROR: Don't know what to do with metric type %s", metric.MetricType)
	}

	p := influxdb2.NewPoint(measurement,
		tags,
		fields,
		time.Now()) //TODO use lastStatusUpdate of metric!

	return writeAPI.WritePoint(context.Background(), p)
}
