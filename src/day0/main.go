package main

import (
	"fmt"
	"time"
)

func main() {
	sum, sub := sumAndSub(5, 3)
	fmt.Println("Sum is", sum, ", sub is ", sub)
	fmt.Println("0K")
	fmt.Println("Time is", time.Now().UTC())

}
func sumAndSub(x int, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}
