package main

import "fmt"

func main() {
	var int1, int2, int3, int4, int5 int
	var char1, char2, char3 byte

	fmt.Scanf("%d %d %d %d %d", &int1, &int2, &int3, &int4, &int5)
	
	var dummy string
	fmt.Scanf("%s", &dummy)
	fmt.Scanf("%c%c%c", &char1, &char2, &char3)

	fmt.Printf("%c%c%c%c%c\n", int1, int2, int3, int4, int5)
	fmt.Printf("%c%c%c\n", char1+1, char2+1, char3+1)
} 