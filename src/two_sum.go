package main

import "fmt"

func main() {
	x := []int{11, 23, 30, 40, 50}
	fmt.Printf("return: ", twoSum(x, 33))
}

func twoSum(nums []int, target int) []int {
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	// indexMap := make(map[int]int)

	// You can also declare and initialize a new map in the same line with this syntax.
	indexMap := map[int]int{}

	// https://gobyexample.com/range
	// range on arrays and slices provides both the index and value for each entry.
	for index, num := range nums {
		// https://gobyexample.com/maps
		// Get a value for a key with name[key].
		/* The optional second return value when getting a value
		 * from a map indicates if the key was present in the map.
		 * true or false
		 */
		if first, found := indexMap[target-num]; found {
			return []int{first, index}
		}
		indexMap[num] = index
	}
	return nil
}
