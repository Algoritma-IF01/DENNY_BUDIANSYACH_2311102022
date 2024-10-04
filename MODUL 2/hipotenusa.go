package main
import "fmt"
func main(){
	var a, b, c float64
	var hipotenusa bool
	
	fmt.Print("masukan a:")
	fmt.Scanln(&a)
	fmt.Print("masukan b:")
	fmt.Scanln(&b)
	fmt.Print("masukan c:")
	fmt.Scanln(&c)
	hipotenusa = (c*c) ==(a*a+b*b)
	fmt.Println("sisi c adalah hipotenusa segitiga a, b, c: ", hipotenusa)
}