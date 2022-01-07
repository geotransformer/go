package main

import "fmt"

var (
	globalInt       int
	globalArr       []int
	globalInterface interface{}
)

func main() {
	// const greeting string = "Hello, World!"
	// var one string = "one"

	// // Only in function body
	// two := "two"

	// _, x := 22, 33
	// var three = 100
	// var four, five = 4, 5
	// four, five = five, four
	// fmt.Println(greeting, one, two, three, four, five, globalInt, x)

	// for index := 0; index < len(greeting); index++ {
	// 	fmt.Printf("%s -- char at %d = %c\n", greeting, index, greeting[index])
	// }

	// m1 := map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	// fmt.Println(m1)
	// m2 := make(map[string]int)
	// for k, v := range m1 {
	// 	m2[v] = k
	// }

	// fmt.Println(m2)
	// a,b := 1,2
	// bar(a,b)
	// fmt.Println(a)

	/*
	15
	30
	*/	
	// closure_func := func_closure(10)
	// fmt.Println(closure_func(5))
	// fmt.Println(closure_func(20))

	/*
	3
	3
	3*/
	// for i := 0; i < 3; i++ {
	// 	defer func ()  {
	// 		fmt.Println(i)
	// 	}()
	// }

	A()
	B()
	C()

}

func A()  {
	fmt.Println("Func A")
}

func B()  {
	defer func ()  {
		if err:=recover(); err != nil {
			fmt.Println("Recover in B")
		}
	}()
	panic("Func B")
}

func C()  {
	fmt.Println("Func c")
}

func foo(a int, b int) (int, string) {
	return 1, "d"
}

func bar(s ...int) {
	s[0] = 8
	fmt.Println(s)
}

// closure
func func_closure(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
