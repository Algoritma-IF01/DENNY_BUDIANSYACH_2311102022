package main

import (
	"fmt"
	"math"
)

func main() {
	var N, x, idx, target int

	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&N)

	arr := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scanln(&arr[i])
	}	
	//print array
	fmt.Println("Isi array:", arr)

	//tampil ganjil
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 0; i < N; i ++ {
		if arr[i]%2 != 0 {
			fmt.Printf("%d ", arr[i])
	
		}
	}
	fmt.Println()

	//tampil genap
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < N; i ++ {
		if arr[i]%2 == 0 {
			fmt.Printf("%d ", arr[i])
	
		}
	}
	fmt.Println()

	//frekensk
	fmt.Print("Masukkan bilangan untuk mencari frekuensi: ")
	fmt.Scanln(&target)
	frequency := 0
	for _, val := range arr {
		if val == target {
			frequency++
		}
	}
	fmt.Printf("Frekuensi bilangan %d dalam array: %d\n", target, frequency)

	//kelipatan x
	fmt.Print("Masukkan bilangan untuk kelipatan indeks: ")
	fmt.Scanln(&x)
	fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
	for i := 0; i < N; i++ {
		if arr[i]%x == 0 {
			fmt.Printf("%d ", arr[i])
		}
	}
	fmt.Println()
	
	//avg
	sum := 0
	for _, val := range arr {
		sum += val
	}
	mean := float64(sum) / float64(len(arr))
	fmt.Printf("Rata-rata dari array: %.2f\n", mean)

	//hapuselemen
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scanln(&idx)
	if idx >= 0 && idx < N {
		arr = append(arr[:idx], arr[idx+1:]...)
		fmt.Println("Isi array setelah penghapusan:", arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	//std
	var std float64
	for _, val := range arr {
		std += math.Pow(float64(val)-mean, 2)
	}
	stdDev := math.Sqrt(std / float64(len(arr)))
	fmt.Printf("Standar deviasi dari array: %.2f\n", stdDev)
}
