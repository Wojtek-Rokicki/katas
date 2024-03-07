package main

/*
Dynamic Programming approach lets you avoid recalculating output for the same input as opposed to recurrence.
*/

// Input Complexity: O(n)
// Memory: O(n)
func climbStairsDynamic(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	f := make([]int, n+1)

	f[1] = 1
	f[2] = 2
	for i := 3; i < n+1; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}

// Input Complexity: O(2^n)
// Memory: O(n) - it is due to stack nature which releases space and inserts new values
func climbStairsRecurrently(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n == 3 {
		return 3
	}
	return climbStairsRecurrently(n-1) + climbStairsRecurrently(n-2)
}
