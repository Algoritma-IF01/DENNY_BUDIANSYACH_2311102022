package main

import (
	"fmt"
)
type arrBalita [100]float64

// berat min & max
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64) {
	*bMin, *bMax = arrBerat[0], arrBerat[0]
	for i := 1; i < len(arrBerat); i++ {
		if arrBerat[i] == 0 { // stop loop untuk data kosong
			break
		}
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

// ratarata berat 
func rataRata(arrBerat arrBalita) float64 {
	var total float64
	var count int

	for i := 0; i < len(arrBerat); i++ {
		if arrBerat[i] == 0 {
			break
		}
		total += arrBerat[i]
		count++
	}
	return total / float64(count)
}

func main() {
	var N int
	var berat arrBalita
	var bMin, bMax float64

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&berat[i])
	}
	hitungMinMax(berat, &bMin, &bMax)
	rerata := rataRata(berat)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rarata berat balita: %.2f kg\n", rerata)
}