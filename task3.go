package main

func MovieDuration(n int) []int {
	data := []int{3, 5, 6, 7, 2}
	results := []int{0, 0}

	for i := 0; i < len(data); i++ {
		for j := i; j < len(data); j++ {
			if data[i]+data[j] == n {
				results[0] = data[i]
				results[1] = data[j]
				return results
			}
		}
	}
	return results
}
