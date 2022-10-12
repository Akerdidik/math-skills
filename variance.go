package math_skills

import "math"

// func Variance(keys []float64, avg float64) (float64, float64) {
// 	var variance float64
// 	for _, val := range keys {
// 		temp := (val - avg)
// 		temp *= temp
// 		variance += temp
// 	}
// 	fl1 := math.Sqrt((float64(variance) / float64(len(keys))))
// 	fl := math.Ceil(float64(float64(variance) / float64(len(keys))))
// 	variance = math.Round(fl)
// 	return variance, fl1
// }
func Stdvar(keys []float64) float64 {
	return stdpop(keys)
}
func stdpop(keys []float64) float64 {
	vs := popvar(keys)
	return math.Sqrt(vs)
}
func Variance(keys []float64) float64 {
	return popvar(keys)
}
func popvar(keys []float64) float64 {
	v := vars(keys)
	return v
}
func vars(keys []float64) float64 {
	avg := Avg(keys)
	var variance float64
	for _, j := range keys {
		variance += (j - avg) * (j - avg)
	}
	return variance / float64(len(keys))
}
