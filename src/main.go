package main

import "fmt"

func main() {

	// Array
	var arr [4]int
	arr[0] = 1
	fmt.Println(arr, len(arr), cap(arr))

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice, len(slice), cap(slice))

	// slice funciona igual que python.
	fmt.Println(slice[0])
	fmt.Println(slice[:3])

	//append
	slice = append(slice, 7)
	fmt.Println(slice)

	// Append dos lista. Anade
	nuevoSlice := []int{8, 9, 10}
	slice = append(slice, nuevoSlice...)
	fmt.Println(slice)

}
