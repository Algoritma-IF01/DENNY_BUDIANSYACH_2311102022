package main

import (
	"fmt"
)

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	fmt.Println("Masukkan karakter satu per satu (akhiri dengan TITIK '.'): ")
	for {
		fmt.Scanf("%c", &input)
		if input == '.' || *n >= NMAX {
			break
		}
		t[*n] = input
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}
func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	fmt.Print("Teks: ")
	cetakArray(tab, m)
	isPalindrom := palindrom(tab, m)
	fmt.Print("Reverse Teks: ")
	balikanArray(&tab, m)
	cetakArray(tab, m)

	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
