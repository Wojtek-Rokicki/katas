package main

func solutionWithCounterMap(s, t string) bool {
	counter := make(map[rune]int)

	for _, x := range s {
		counter[x] += 1
	}

	for _, x := range t {
		counter[x] -= 1
	}

	for _, v := range counter {
		if v != 0 {
			return false
		}
	}

	return true
}

func solutionWithCounterArray(p, q string) bool {
	counter := [26]int{}
	length := len(q)
	if length != len(p) {
		return false
	}
	min := int('a')
	for _, v := range p {
		index := int(v) - min
		counter[index] += 1

	}
	for _, v := range q {
		index := int(v) - min
		counter[index] -= 1
	}
	for i := 0; i < 26; i++ {
		if counter[i] != 0 {
			return false
		}
	}
	return true

}

func main() {

}
