package main
import "fmt"

func f(x int) int {
	return x * x
}
func g(x int) int {
	return x - 2
}
func h(x int) int {
	return x + 1
}

func main() {
	var a, b, c int 
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&a, &b, &c)

	fungsi1 := f(g(h(a)))
	fungsi2 := g(h(f(b)))
	fungsi3 := h(f(g(c)))

	fmt.Printf("(fogoh) (%d)= %d\n", a, fungsi1)
	fmt.Printf("(gohof) (%d)= %d\n", b, fungsi2)
	fmt.Printf("(hofog) (%d)= %d\n", c, fungsi3)
}