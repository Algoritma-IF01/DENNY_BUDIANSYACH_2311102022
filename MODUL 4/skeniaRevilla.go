package main

import (
	"fmt"
)

func cetakderet(n int) {
	for n >= 1 {
		fmt.Print(n, " ")
		if n == 1 {
			break
		}
		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanf("%d", &n)
	cetakderet(n)
}
