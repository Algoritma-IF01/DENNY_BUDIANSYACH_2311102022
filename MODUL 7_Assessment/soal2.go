package main

import "fmt"

func main() {
	var rombongan, summenu, oragn int
	var sisa string

	const hargabiasa = 10000
	const hargapermenu = 2500
	const fiftymenu = 100000

	fmt.Print("Masukan jumlah rombongan: ")
	fmt.Scan(&rombongan)

	for i := 1; i <= rombongan; i++ {
		fmt.Print("Jumlah menu, jumlah orang, sisa makan (y/t): ")
		fmt.Scan(&summenu, &oragn, &sisa)

		totalharga := 0

		if summenu <= 3 {
			totalharga = hargabiasa
		} else if summenu > 3 && summenu <= 50 {
			totalharga = hargabiasa + (summenu-3)*hargapermenu
		} else {
			totalharga = fiftymenu
		}

			if sisa == "y" {
			totalharga *= oragn
		}

		fmt.Printf("Total biaya rombongan %d: %d\n", i, totalharga)
	}
}
