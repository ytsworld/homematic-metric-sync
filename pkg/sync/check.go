package sync

import "log"

/**
	Returns false if more then the tolerated percentage for Temperature sensors + Energy Consumption monitors return zero values (indicating an error at Homematic IP Cloud)
**/
func SanityCheck(metrics *Metrics, sanityCheckPercentage int16) bool {

	var total, zeroValues int16 = 0, 0

	for _, m := range metrics.Metrics {
		total++

		if m.MetricType == "Climate" {
			if m.Climate.Temperature == 0 {
				zeroValues++
			}
		} else if m.MetricType == "Energy" {
			if m.Energy.TotalCounter == 0 {
				zeroValues++
			}
		}
	}

	maxZeroValues := float64(total) / 100.0 * float64(sanityCheckPercentage)

	log.Printf("%d out of %d metrics have zero values (%f allowed)", zeroValues, total, maxZeroValues)

	return float64(zeroValues) <= maxZeroValues
}
