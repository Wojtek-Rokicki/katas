package main

import "fmt"

func solutionWithHashMap(nums []int) bool {
	uniqueMap := make(map[int]bool)
	for _, elem := range nums {
		_, ok := uniqueMap[elem]
		if ok {
			return true
		}
		uniqueMap[elem] = true
	}
	return false
}

func solutionWithArray(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	var min, max int = nums[0], nums[0]
	// Get min max with O(n)
	for _, element := range nums {
		if element < min {
			min = element
		}
		if element > max {
			max = element
		}
	}

	t := make([]bool, max-min+1)

	for _, element := range nums {
		elementIndex := element - min
		if t[elementIndex] {
			return true
		}
		t[elementIndex] = true
	}

	return false
}

func main() {
	fmt.Println("Hello world")
}
