package math_skills

import (
	"sort"
)

func Median(keys []float64) float64 {
	sort.Float64s(keys)
	var median float64
	if len(keys)%2 == 0 {
		median = Avg(keys[len(keys)/2-1 : len(keys)/2+1])
	} else {
		median = keys[len(keys)/2]
	}
	return median
}
