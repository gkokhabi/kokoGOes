/*
You are given an integer array cost where cost[i] is the cost of ith step on a staircase. 
Once you pay the cost, you can either climb one or two steps.
You can either start from the step with index 0, or the step with index 1.
Return the minimum cost to reach the top of the floor.
*/

func minCostClimbingStairs(cost []int) int {
	a, b := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
			if (b < a) {
					a, b = b, cost[i] + b
			} else {
					a, b = b, cost[i] + a
			}
	}
	
	if (a < b) {
			return a
	}
	return b
}
