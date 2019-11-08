package unfairness_metric

type Metric struct {
	First float64
	Total float64
}

func New() *Metric {
	return new(Metric)
}

func ConvertToXYRange(metricList []*Metric) ([]float64, []float64) {
	xRange := make([]float64, 0)
	yRange := make([]float64, 0)

	for i, m := range metricList {
		xRange = append(xRange, float64(i+1))
		yRange = append(yRange, m.First/m.Total)
	}
	return xRange, yRange
}
