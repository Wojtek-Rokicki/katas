package main

func maxSlidingWindow(nums []int, k int) []int {
	deque := make([]int, 0)
	result := make([]int, 0)
	for i := range nums {
		// drop all lower elements
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] { // pop last element if less < actual element
			// pop last element
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)

		// check if max element is out of window
		if i-deque[0] >= k {
			deque = deque[1:]
		}

		// im inside the window
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}

	}
	return result
}
