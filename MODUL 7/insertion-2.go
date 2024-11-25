package main

import "fmt"

const nMax = 7919
type Buku struct {
    id, judul, penulis, penerbit string
    eksemplar, tahun, rating int
}

type DaftarBuku struct {
    pustaka  []Buku
    nPustaka int
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
    for i := 0; i < n; i++ {
        var buku Buku
        fmt.Println("\nData buku ke-", i+1)
        fmt.Print("ID: ")
        fmt.Scan(&buku.id)
        fmt.Print("Judul: ")
        fmt.Scan(&buku.judul)
        fmt.Print("Penulis: ")
        fmt.Scan(&buku.penulis)
        fmt.Print("Penerbit: ")
        fmt.Scan(&buku.penerbit)
        fmt.Print("Eksemplar: ")
        fmt.Scan(&buku.eksemplar)
        fmt.Print("Tahun: ")
        fmt.Scan(&buku.tahun)
        fmt.Print("Rating: ")
        fmt.Scan(&buku.rating)
        
        pustaka.pustaka = append(pustaka.pustaka, buku)
    }
    pustaka.nPustaka = n
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
    if n == 0 {
        fmt.Println("Tidak ada buku")
        return
    }
    
    maxRating := pustaka.pustaka[0]
    for i := 1; i < n; i++ {
        if pustaka.pustaka[i].rating > maxRating.rating {
            maxRating = pustaka.pustaka[i]
        }
    }
    
    fmt.Println("\nBuku Terfavorit:")
    fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\nRating: %d\n",
        maxRating.judul, maxRating.penulis, maxRating.penerbit, maxRating.tahun, maxRating.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
    for i := 1; i < n; i++ {
        key := pustaka.pustaka[i]
        j := i - 1
        
        for j >= 0 && pustaka.pustaka[j].rating < key.rating {
            pustaka.pustaka[j+1] = pustaka.pustaka[j]
            j--
        }
        pustaka.pustaka[j+1] = key
    }
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
    fmt.Println("\n5 Buku Rating Tertinggi:")
    limit := 5
    if n < 5 {
        limit = n
    }
    
    for i := 0; i < limit; i++ {
        fmt.Printf("%d. %s (Rating: %d)\n", 
            i+1, pustaka.pustaka[i].judul, pustaka.pustaka[i].rating)
    }
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
    fmt.Printf("\nBuku dengan rating %d:\n", r)
    
    left := 0
    right := n - 1
    found := false
    
    for left <= right {
        mid := (left + right) / 2
        
        if pustaka.pustaka[mid].rating == r {
            found = true
            fmt.Printf("ID: %s\nJudul: %s\nPenulis: %s\nPenerbit: %s\n"+
                "Eksemplar: %d\nTahun: %d\nRating: %d\n",
                pustaka.pustaka[mid].id, pustaka.pustaka[mid].judul,
                pustaka.pustaka[mid].penulis, pustaka.pustaka[mid].penerbit,
                pustaka.pustaka[mid].eksemplar, pustaka.pustaka[mid].tahun,
                pustaka.pustaka[mid].rating)
            break
        }
        
        if pustaka.pustaka[mid].rating < r {
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    
    if !found {
        fmt.Println("Tidak ditemukan buku dengan rating %d:", r)
    }
}

func main() {
    var n, ratingCari int
    var pustaka DaftarBuku
    
    fmt.Print("Jumlah buku: ")
    fmt.Scan(&n)
    
    DaftarkanBuku(&pustaka, n)
    
    CetakTerfavorit(pustaka, n)
    
    UrutBuku(&pustaka, n)
    
    Cetak5Terbaru(pustaka, n)
    
    fmt.Print("\nRating buku yang dicari: ")
    fmt.Scan(&ratingCari)
    CariBuku(pustaka, n, ratingCari)
}