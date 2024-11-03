package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}
		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorA < skorB {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}
	fmt.Println("Pertandingan selesai")
}
