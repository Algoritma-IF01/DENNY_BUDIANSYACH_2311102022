package main
import "fmt"

func INSRTsort(arr []int) {
    for i := 1; i < len(arr); i++ {
        key := arr[i]
        j := i - 1
        for j >= 0 && arr[j] > key {
            arr[j+1] = arr[j]
            j--
        }
        arr[j+1] = key
    }
}

func main() {
    for {
        var bilangan int
        arr := []int{}

        for {
            fmt.Scan(&bilangan)
            if bilangan < 0 {
                break
            }
            arr = append(arr, bilangan)
        }

        if len(arr) == 0 {
            break
        }
        INSRTsort(arr)
        fmt.Println("Hasil urut:", arr)

        jarakTetap := true
        jarak := 0
        if len(arr) > 1 {
            jarak = arr[1] - arr[0]
            for i := 1; i < len(arr)-1; i++ {
                if arr[i+1]-arr[i] != jarak {
                    jarakTetap = false
                    break
                }
            }
        }

        if jarakTetap && len(arr) > 1 {
            fmt.Printf("Data berjarak %d\n", jarak)
        } else {
            fmt.Println("Data berjarak tidak tetap")
        }
    }
}
