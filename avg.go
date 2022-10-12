package math_skills

func Avg(keys []float64) float64 {
	var sum float64
	for _, j := range keys {
		sum += j
	}
	return sum / float64(len(keys))
}
