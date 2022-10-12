package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"math_skills"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		return
	}
	file, _ := ioutil.ReadFile(args[0])
	str := string(file)
	res := strings.Fields(str)
	keys := make([]float64, 0)
	for _, val := range res {
		key, err := strconv.Atoi(val)
		if err != nil {
			return
		}
		keys = append(keys, float64(key))
	}
	//fmt.Println(variance)
	fmt.Println("Average:", math.Round(math_skills.Avg(keys)))
	fmt.Println("Median:", math.Round(math_skills.Median(keys)))
	fmt.Println("Variance:", int(math.Round(math_skills.Variance(keys))))
	fmt.Println("Standard Deviation:", math.Round(math_skills.Stdvar(keys)))
}
