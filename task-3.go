package main

func findMovieDuration(data []int, n int) (int, int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i; j < len(data); j++ {
			if data[i]+data[j] == n {
				return data[i], data[j]
			}
		}
	}

	return 0, 0
} //slice

// func findMovieDuration(data map[string]int, n int) (string, int, string, int) {
// 	for movie1, duration1 := range data {
// 		for movie2, duration2 := range data {
// 			if movie1 != movie2 && duration1+duration2 == n {
// 				return movie1, duration1, movie2, duration2
// 			}
// 		}
// 	}

// 	return "", 0, "", 0
// } //map
