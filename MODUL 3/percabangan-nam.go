//1. program akan error karena terdap beberapa kesalah dalam algoritma program
//2. - 	kesalahan penggunaan perulangan if. seharurnya program menggunakan satu percabangan if yang berkelanjutan.
//		program pada modul menggunakan percabangan if baru untuk setiap kondisi, hal ini membuat program hanya
//		melakukan pengecekan hingga kondisi if terbawah alih alih mengklasifikasikan nilai nam
//   -	kesalahan penggunaan variabel nam pada pengkondisian. seharusnya variabel yang digunakan adalah nmk yang 
//		bertipe data string
package main

import "fmt"

func main() {
	var nam float64
	var nmk string
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "AB"
	} else if nam > 65 {
		nmk = "B"
	} else if nam > 57.5 {
		nmk = "BC"
	} else if nam > 50 {
		nmk = "C"
	} else if nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}
	fmt.Println("Nilai mata kuliah: ", nmk)
}