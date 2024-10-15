package main

import "fmt"

func main() {
    var berat, kg, gr int
    var biayaPerKg, biayaTambahan, totalBiaya int

    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

	kg= berat/1000
    gr= berat%1000
	biayaPerKg = kg*10000

    if gr>=500 {
        biayaTambahan= gr*5
    }else{
		biayaTambahan=gr*15
	}

	if kg>=10{
		totalBiaya = biayaPerKg 
	}else{
		totalBiaya = biayaPerKg+biayaTambahan
	}

    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gr)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaTambahan)
    fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
