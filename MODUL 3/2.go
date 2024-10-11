package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungacount int

	for{
		fmt.Printf("Bunga %d: ", bungacount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input)=="selesai"{
			break
		}
		if pita == ""{
			pita = input
		}else{
			pita += " - "+input
		}
		bungacount++
	}
	fmt.Printf("pita: %s\n", pita)
	fmt.Printf("bunga: %d\n", bungacount)
}