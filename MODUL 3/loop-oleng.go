package main

import (
	"fmt"
)

func main() {
	var kiri, kanan float32

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kiri, &kanan)
		if kiri+kanan >= 151 {
			break
		}
		selisih := kiri - kanan
		if selisih < 0 {
			selisih = selisih*(-1)
		}
		oleng := selisih > 9
		fmt.Println("Sepeda motor Pak Andi akan oleng:", oleng)
	}
	fmt.Println("Sepeda motor Pak Andi akan oleng: TRUE")
	fmt.Println("Proses selesai")
}
