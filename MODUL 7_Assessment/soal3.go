package main

import "fmt"

func kelipatn4() int {
	var bilangan int
	fmt.Scan(&bilangan)

	if bilangan < 0 {
		return 0
	}

	if bilangan%4 == 0 {
		return bilangan + kelipatn4()
	}

	return kelipatn4()
}

func main() {
	fmt.Println("Masukkan bilangan (akhiri dengan bilangan negatif):")

	hasil := kelipatn4()

	fmt.Println("Jumlah bilangan kelipatan 4 =", hasil)
}
