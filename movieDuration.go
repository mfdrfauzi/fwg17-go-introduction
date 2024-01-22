package main

import "fmt"

func findMovieDuration(data []int, n int) (int, int) {
	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i]+data[j] == n {
				return data[i], data[j]
			}
		}
	}

	return 0, 0
} //slice

// func findMovieDuration(data map[string]int, n int) (string, string) {
// 	for movie1, duration1 := range data {
// 		for movie2, duration2 := range data {
// 			if movie1 != movie2 && duration1+duration2 == n {
// 				return movie1, movie2
// 			}
// 		}
// 	}

// 	return "", ""
// } //map

func main() {
	data := []int{3, 5, 6, 7, 2} //slice

	// data := map[string]int{
	// 	"Fast Furious": 5,
	// 	"Tom & Jerry":  3,
	// 	"Avengers":     6,
	// 	"Interstellar": 7,
	// 	"Larva":        2,
	// } //map

	var n int

	fmt.Print("Masukkan durasi penerbangan: ")
	fmt.Scan(&n)

	result1, result2 := findMovieDuration(data, n)
	if result1 != 0 && result2 != 0 { //slice
		// if result1 != "" && result2 != "" { //map
		fmt.Printf("Hasil penjumlahan durasi film: %d dan %d\n", result1, result2) //slice
		// fmt.Printf("Hasil penjumlahan durasi film: %#v dan %#v\n", result1, result2) //map
	} else {
		fmt.Println("Tidak ada koleksi film yang cocok dengan durasi penerbangan")
	}
}
