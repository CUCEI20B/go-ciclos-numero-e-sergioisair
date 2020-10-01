package main

import "fmt"

func main() {
	var n int
	var e float64 = 0
	var x float64 = 1
	fmt.Scan(&n)
	for i := 0; i <= n; i++ {
		x = 1
		for j := i; j >= 1; j-- {
			x = x * float64(j)
		}
		e = e + (1 / x)
	}
	fmt.Print(e)
}
