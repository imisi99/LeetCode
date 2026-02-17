package main

import (
	"fmt"
	"sort"
)

// Fracknapsack -> Time 0(nlogn) Space 0(n)
func Fracknapsack(knapsack int, items [][]int) float64 {
	var totalValue float64
	ratios := make([][]float64, len(items))

	for i, item := range items {
		weight, value := item[0], item[1]
		ratio := float64(value) / float64(weight)
		ratios[i] = []float64{float64(i), ratio}
	}

	sort.Slice(ratios, func(i, j int) bool {
		return ratios[i][1] > ratios[j][1]
	})

	fmt.Println(ratios)

	weight := knapsack
	for _, ratio := range ratios {
		i := ratio[0]
		currWeight, currVal := items[int(i)][0], items[int(i)][1]
		if currWeight <= weight {
			totalValue += float64(currVal)
			weight -= currWeight
		} else {
			frac := float64(weight) / float64(currWeight)
			fmt.Println(frac, weight)
			totalValue += (frac * float64(currVal))
			break
		}
	}

	return totalValue
}

func main() {
	items := [][]int{
		{10, 70}, {5, 50}, {7, 20}, {4, 50},
	}
	fmt.Println(Fracknapsack(20, items))
}
