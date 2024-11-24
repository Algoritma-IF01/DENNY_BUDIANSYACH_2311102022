package main

import "fmt"

func SLCTsort(arr []int) {
    for i := 0; i < len(arr)-1; i++ {
        mindex := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[mindex] {
                mindex = j
            }
        }
        arr[i], arr[mindex] = arr[mindex], arr[i]
    }
}

func ganjlgnap(arr []int) {
    for _, num := range arr {
        if num%2 == 1 {
            fmt.Printf("%d ", num)
        }
    }
    for _, num := range arr {
        if num%2 == 0 {
            fmt.Printf("%d ", num)
        }
    }
    fmt.Println()
}

func main() {
    var jumlahDaerah int
    fmt.Print("jumlah daerah: ")
    fmt.Scan(&jumlahDaerah)

    for i := 0; i < jumlahDaerah; i++ {
        var jumlahRumah int
        fmt.Printf("jumlah rumah kerabat untuk daerah %d: ", i+1)
        fmt.Scan(&jumlahRumah)

        rumah := make([]int, jumlahRumah)
        
        fmt.Printf("Nomor rumah untuk daerah %d: ", i+1)
        for j := 0; j < jumlahRumah; j++ {
            fmt.Scan(&rumah[j])
        }

        SLCTsort(rumah)

        fmt.Printf("Hasil daerah %d: ", i+1)
        ganjlgnap(rumah)
    }
}