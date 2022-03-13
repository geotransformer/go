package main

import (
	"fmt"
	//"time"
)
func main() {
	// go go_thread()
	// go go_thread()
	// go go_thread()
	// go go_thread()
	// time.Sleep(2 * time.Second)
	// fmt.Println("main thread")

	c := make(chan bool)
	go func ()  {
		fmt.Println("go go go!!!")
		c <- true
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}

}

// func go_thread(){
// 	fmt.Println("go thread")
// }
