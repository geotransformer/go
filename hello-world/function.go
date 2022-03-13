package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}
func main() {
	a, b := swap("Apple", "Ball")
	fmt.Println(a, b)
}
