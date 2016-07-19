package main

import (
	"fmt"
)

func main() {

	println("Loop an array: ")

	arrayType := [5]int{1, 2, 3, 4, 5}

	for _, value := range arrayType {
		println(value)
	}

	println("\nLoop a slice: ")

	sliceType := make([]int, 0)

	sliceType = append(sliceType, 1)
	sliceType = append(sliceType, 2)
	sliceType = append(sliceType, 3)
	sliceType = append(sliceType, 4)
	sliceType = append(sliceType, 5)

	for _, value := range sliceType {
		println(value)
	}

	println("\nLoop a map: ")

	mapType := make(map[int]int)

	mapType[1] = 1
	mapType[2] = 2
	mapType[3] = 3
	mapType[4] = 4
	mapType[5] = 5

	for key, value := range mapType {
		fmt.Printf("Key: %d, Value: %d \n", key, value)
	}
}
