package main
import "fmt"
func main(){
	var tahun int
	
	fmt.Print("masukan tahun:")
	fmt.Scanln(&tahun)

	if tahun%4==0{
		fmt.Println("tahum kabisat : True")
	}else{
		fmt.Println("tahum kabisat : False")
	}
} 