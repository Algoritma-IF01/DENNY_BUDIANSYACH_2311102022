package main

import "fmt"

func SLCsort(arr []int) []int {
    for i := 0; i < len(arr)-1; i++ {
        min := i
        for j := i + 1; j < len(arr); j++ {
            if arr[j] < arr[min] {
                min = j
            }
        }
        arr[i], arr[min] = arr[min], arr[i]
    }
    return arr
}

func main() {
    var daerah, rumah int

    fmt.Print("jumlah daerah: ")
    fmt.Scan(&daerah)

    for i := 1; i <= daerah; i++ {
        fmt.Printf("jumlah rumah kerabat untuk daerah %d: ", i)
        fmt.Scan(&rumah)

        arr := make([]int, rumah)
        fmt.Printf("Nomor rumah untuk daerah %d: ", i)
        for j := 0; j < rumah; j++ {
            fmt.Scan(&arr[j])
        }

        sorted := SLCsort(arr)

        fmt.Printf("daerah %d: %v\n", i, sorted)
    }
}
