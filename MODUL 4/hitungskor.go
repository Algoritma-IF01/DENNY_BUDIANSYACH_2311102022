package main

import "fmt"

func hitungSkor(soal []int) (int, int) {
    skor, totalWaktu := 0, 0
    for _, waktu := range soal {
        if waktu != 301 {
            skor++
            totalWaktu += waktu
        }
    }
    return skor, totalWaktu
}

func main() {
    var nama1, nama2 string
    var soal1, soal2 [8]int

    fmt.Scan(&nama1)
    for i := 0; i < 8; i++ {
        fmt.Scan(&soal1[i])
    }

    fmt.Scan(&nama2)
    for i := 0; i < 8; i++ {
        fmt.Scan(&soal2[i])
    }

    skor1, waktuTotal1 := hitungSkor(soal1[:])
    skor2, waktuTotal2 := hitungSkor(soal2[:])

    if skor1 > skor2 || (skor1 == skor2 && waktuTotal1 < waktuTotal2) {
        fmt.Printf("%s %d %d\n", nama1, skor1, waktuTotal1)
    } else {
        fmt.Printf("%s %d %d\n", nama2, skor2, waktuTotal2)
    }
}
