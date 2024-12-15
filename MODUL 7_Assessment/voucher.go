package main

import (
	"fmt"
)

func validvoucher(voucher string) bool {
	length := len(voucher)

	if length != 5 && length != 6 {
		return false
	}

	firstDigit := int(voucher[0] - '0')
	secondDigit := int(voucher[1] - '0')
	lastDigit := int(voucher[length-1] - '0')
	secondLastDigit := int(voucher[length-2] - '0')

	if firstDigit*secondDigit != secondLastDigit*lastDigit {
		return false
	}
	if length == 5 {
		midDigit := int(voucher[2] - '0')
		if midDigit%2 != 0 {
			return false
		}
	} else if length == 6 {
		midDigit1 := int(voucher[2] - '0')
		midDigit2 := int(voucher[3] - '0')
		midDigit3 := int(voucher[4] - '0')

		if midDigit1%2 != 0 || midDigit2%2 != 0 || midDigit3%2 != 0 {
			return false
		}
	}

	return true
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah mahasiswa: ")
	fmt.Scan(&N)

	validCount := 0
	invalidCount := 0

	for i := 0; i < N; i++ {
		var voucher string
		fmt.Print("Masukkan nomor seri voucher: ")
		fmt.Scan(&voucher)

		if validvoucher(voucher) {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Println("Banyak voucher valid:", validCount)
	fmt.Println("Banyak voucher tidak valid:", invalidCount)
}
