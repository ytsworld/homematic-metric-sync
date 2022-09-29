package sync

import "log"

const SanityCheckPercentage = 20

/**
	Returns false if more then the tolerated percentage for Temperature sensors + Energy Consumption monitors return zero values (indicating an error at Homematic IP Cloud)
**/
func SanityCheck(metrics *Metrics) bool {

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

	log.Printf("%d out of %d metrics have zero values", zeroValues, total)

	return zeroValues <= total/100*SanityCheckPercentage
}
