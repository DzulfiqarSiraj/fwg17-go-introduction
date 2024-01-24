package main

func MovieDuration(n int) (int, int) {
	data := []int{3, 5, 6, 7, 2}

	for i := 0; i < len(data)-1; i++ {
		for j := i; j < len(data); j++ {
			if data[i]+data[j] == n {
				return data[i], data[j]
			}
		}
	}
	return 0, 0
}
