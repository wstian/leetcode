package main

import "sort"

func twoCitySchedCost(costs [][]int) int {
	sort.SliceStable(costs, func(i, j int) bool {
		return (costs[i][0] - costs[i][1]) < (costs[j][0] - costs[j][1])
	})
	total := 0
	half := len(costs) / 2
	for i, cost := range(costs)  {
		if i < half {
			total += cost[0]
		} else {
			total += cost[1]
		}
	}
	return total
}
