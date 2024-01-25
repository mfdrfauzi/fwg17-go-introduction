package main

import (
	"fmt"
	"strings"
)

func main() {
	//Print Segitiga

	fmt.Println("---Print Segitiga Terbalik---")
	printSegitiga(12)

	//Generate Password
	fmt.Println("---Generate Password---")
	var password, level string

	fmt.Print("Enter password: ")
	fmt.Scan(&password)

	for {
		fmt.Print("Enter level (low, med, strong): ")
		fmt.Scan(&level)

		if strings.ToLower(level) == "low" || strings.ToLower(level) == "med" || strings.ToLower(level) == "strong" {
			break
		} else {
			fmt.Println("Invalid level. Please enter 'low', 'med', or 'strong'.")
		}
	}

	result := genPass(password, level)
	if result != "" {
		fmt.Println(result) //perlu / tidak vlidasi di sini
	}

	//Movie Duration
	fmt.Println("---Find Movie---")
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
