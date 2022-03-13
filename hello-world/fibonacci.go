package main

import "fmt"

func main() {
	fmt.Println(fib(6))

}
func fib(n int) int {
	// Sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610
	// 1. recursive
	// 2. DP
	// 3. keep updating f(n-1), f(n-2)
	if n < 2 {
		return n
	}
	var f0, fi, f1, i = 0, 0, 1, 2
	for i <= n {
		fi = f0 + f1
		f0, f1 = f1, fi
		i++
	}
	return fi
}
