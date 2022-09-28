package sync

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

	return zeroValues > total/100*SanityCheckPercentage
}
