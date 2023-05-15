package sync

import (
	"testing"
)

func TestSanityCheck(t *testing.T) {
	m := Metrics{}
	m.Metrics = append(m.Metrics, &Metric{
		DeviceId:   "abc",
		MetricType: "Climate",
		Climate: &ClimateMetric{
			Temperature: 0,
		},
	})

	sanity := SanityCheck(&m, 50)

	if sanity {
		t.Fatalf("Unexpected: First Sanity returned true")
	}

	m.Metrics = append(m.Metrics, &Metric{
		DeviceId:   "abc",
		MetricType: "Energy",
		Energy: &EnergyMetric{
			TotalCounter: 1,
		},
	})

	sanity2 := SanityCheck(&m, 50)

	if !sanity2 {
		t.Fatalf("Unexpected: Second Sanity returned false")
	}

	sanity3 := SanityCheck(&m, 20)

	if sanity3 {
		t.Fatalf("Unexpected: Third Sanity returned true")
	}
}
