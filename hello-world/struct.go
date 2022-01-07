package main

import "fmt"

// type <Name> struct {}

type Person struct {
	name string
	age  int
}

type City struct {
	name string
}

func main() {
	// a := &Person{
	// 	name: "Jack",
	// 	age: 66,
	// }
	// // a.name = "Tom"
	// // a.age = 23
	// fmt.Println(a)
	// foo(a)
	// fmt.Println(a)

	a := &struct{
		name string
		age int
		add City
	}{
		name: "Tom",
		age: 23,
		add: City{
			name: "Ott",
		},
	}

	fmt.Println(a)


}

// func foo(p *Person){
// 	// no need *
// 	p.name = "Change"
// 	fmt.Println("after", p)
// }
