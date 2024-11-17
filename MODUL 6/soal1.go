package main
import("fmt")

func main(){
	var N int
	fmt.Print("Masukan jumlah anak kelinci: ")
	fmt.Scan(&N)

	if N<=0 || N>1000{
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}
	weights := make([]float64,N)
	fmt.Println("Masukan berat anak kelinci: ")
	for i :=0; i<N; i++{
		fmt.Scan(&weights[i])
	}
	minweight, maxweight := weights[0], weights[0]

	for _, weight:= range weights[1:]{
		if weight<minweight{
			minweight = weight
		}
		if weight> maxweight{
			maxweight = weight
		}
	}
	fmt.Printf("Berat kelinci terkecil: %.2f\n", minweight)
	fmt.Printf("Berat kelinci terbesar: %.2f\n", maxweight)
}