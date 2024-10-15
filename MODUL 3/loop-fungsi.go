package main

import (
	"fmt"
)

func akar2(k int) float64 {
	product := 1.0
	for i := 0; i <= k; i++ {
		pembilang := (4*float64(i) + 2) * (4*float64(i) + 2)
		penyebut := (4*float64(i) + 1) * (4*float64(i) + 3)
		product *= pembilang / penyebut
	}
	return product
}

func main() {
	var K int
	fmt.Print("Nilai K = ")
	fmt.Scanln(&K)

	result := akar2(K)
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
