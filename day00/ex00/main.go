package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"sort"
)

var meanf bool
var medianf bool
var modef bool
var sdf bool

func init() {
	flag.BoolVar(&meanf, "mean", false, "mean flag work")
	flag.BoolVar(&medianf, "median", false, "median flag work")
	flag.BoolVar(&modef, "mode", false, "mode flag work")
	flag.BoolVar(&sdf, "sd", false, "sd flag work")
	flag.Parse()
	if !meanf && !medianf && !modef && !sdf {
		meanf = true
		medianf = true
		modef = true
		sdf = true
	}
}

func mean(slice []int) float64 {
	var sum float64
	for _, v := range slice {
		sum += float64(v)
	}
	return sum / float64(len(slice))
}

func median(slice []int) float64 {
	if len(slice)%2 == 1 {
		return float64(slice[len(slice)/2])
	} else {
		return mean([]int{slice[len(slice)/2-1], slice[len(slice)/2]})
	}
}

func mode(slice []int) int {
	var curCount int = 0
	var resKey int
	var dic = make(map[int]int)
	for _, v := range slice {
		dic[v]++
	}
	for key, count := range dic {
		if count > curCount {
			curCount = count
			resKey = key
		} else if count == curCount {
			if key < resKey {
				resKey = key
			}
		}
	}
	return resKey
}

func sd(slice []int) float64 {
	var sum float64 = 0
	var avg float64 = mean(slice)
	for _, v := range slice {
		sum += math.Pow(float64(v)-avg, 2)
	}
	avg = sum / float64(len(slice))
	return math.Sqrt(avg)
}

func main() {
	var slice = make([]int, 0, 20)
	var tmp int
	var err error
	fmt.Println("hello")

	for _, err = fmt.Scanln(&tmp); err != io.EOF; _, err = fmt.Scanln(&tmp) {
		if tmp < -100000 || tmp > 100000 {
			fmt.Println("Number out of range.")
		} else if err != nil {
			fmt.Println("Error: Wrong input,", err)
		} else {
			slice = append(slice, tmp)
		}
	}
	
	sort.Ints(slice)
	if len(slice) > 0 {
		if meanf {
			fmt.Printf("Mean: %.2f\n", mean(slice))
		}
		if medianf {
			fmt.Printf("Median: %.2f\n", median(slice))
		}
		if modef {
			fmt.Printf("Mode: %d\n", mode(slice))
		}
		if sdf {
			fmt.Printf("SD: %.2f\n", sd(slice))
		}
	}
}
