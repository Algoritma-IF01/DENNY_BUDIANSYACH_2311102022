package main

import "fmt"

func fakctorial(n int)int{
	if n == 0 || n==1{
		return 1
	}
	result := 1
	for i := 2; i <= n; i++{
		result *= i
	}
	return result
}

func permutation (n, r int)int{
	return fakctorial(n)/ fakctorial(n-r)
}

func combination(n, r int)int{
	return fakctorial(n)/ (fakctorial(r)* fakctorial(n-r))
}
func main() {
	var a, b, c, d int
	fmt.Println ("masukan nilai a, b, c, d:")
	fmt.Scan(&a, &b, &c, &d)

	p1 := permutation(a,c)
	c1 := combination(a,c)
	p2 := permutation(b,d)
	c2 := combination(b,d)

	fmt.Printf("P(%d,%d) = %d\n", a,c,p1)
	fmt.Printf("C(%d,%d) = %d\n", a,c,c1)
	fmt.Printf("P(%d,%d) = %d\n", b,d,p2)
	fmt.Printf("C(%d,%d) = %d\n", b,d,c2)

	
}