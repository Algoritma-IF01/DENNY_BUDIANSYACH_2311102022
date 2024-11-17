package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Println("Masukkan jumlah ikan dan kapasitas wadah: ")
	fmt.Scan(&x, &y)

	if x <= 0 || y <= 0 || x > 1000 {
		fmt.Println("Jumlah ikan atau kapasitas wadah tidak valid")
		return
	}
	berat := make([]float64, x)
	fmt.Println("Masukkan berat ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	var currentWadahWeight float64
	beratwadah := []float64{}
	for i, weight := range berat {
		currentWadahWeight += weight
		if (i+1)%y == 0 || i == x-1 {
			beratwadah = append(beratwadah, currentWadahWeight)
			currentWadahWeight = 0
		}
	}

	var totalWeight float64
	for _, w := range beratwadah {
		totalWeight += w
	}
	averageWeight := totalWeight / float64(len(beratwadah))

	// Output hasil
	fmt.Println("Total berat setiap wadah: ")
	for _, w := range beratwadah {
		fmt.Printf("%.2f ", w)
	}
	fmt.Println()
	fmt.Printf("Berat rata-rata setiap wadah: %.2f\n", averageWeight)
}
