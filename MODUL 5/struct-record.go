package main
import ("fmt")

type waktu struct{
	jam, menit, detik int
}

func main(){
	var wparkir, wpulang, durasi waktu
	var dparkir, dpulang, lparkir int
	fmt.Scan(&wparkir.jam, &wparkir.menit, &wparkir.detik)
	fmt.Scan(&wpulang.jam, &wpulang.menit, &wpulang.detik)
	dparkir = wparkir.detik+wparkir.menit*60+wparkir.jam*3600
	dpulang = wpulang.detik+wpulang.menit*60+wpulang.jam*3600
	lparkir = dpulang-dparkir
	durasi.jam = lparkir/3600
	durasi.menit = lparkir%3600/60
	durasi.detik = lparkir %3600%60
	fmt.Printf("lama parkir : %d jam %d menit %d detik", durasi.jam, durasi.menit, durasi.detik)
}