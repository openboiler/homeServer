package util

type WeightedAvgElement struct {
	Weight float32
	Value  float32
}

func WeightedAvg(elements []WeightedAvgElement) float32 {
	var numerator float32
	var denomominator float32
	for _, e := range elements {
		numerator += (e.Value * e.Weight)
		denomominator += e.Weight
	}
	return numerator / denomominator
}
