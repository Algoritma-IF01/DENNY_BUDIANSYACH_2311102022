package main

import "fmt"

func main() {
	var celcius float64
	fmt.Print("Temperatur Celcius: ")
	fmt.Scanln(&celcius)

	fmt.Println("Derajat Reamur:", celcius*4/5)
	fmt.Println("Derajat Fahrenheit:", celcius*9/5+32)
	fmt.Println("Derajat Kelvin:", celcius+273.15)
}
