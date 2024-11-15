package main

import "fmt"

//array is of fized size and memory is allocated at compile time, something very different to slices
// func getFixedArrayOfGivenSize() [9]int {
// 	const arraySize = 9
// 	arr := [arraySize]int{}

// 	return arr
// }

// //slices are built on top of arrays and memory of slices is dynamically managed
// func getSlice() []int {
// 	return []int{23, 321, 233, 13}
// }

func mutateCopiedArray(arr [1]int) {
	arr[0] = 25
}

func mutateOriginalArray(arr *[1]int) {
	(*arr)[0] = 25
}

func main() {
	// arr := getFixedArrayOfGivenSize()
	// fmt.Println(arr)

	// slice := getSlice()
	// fmt.Println(slice)
	arr := [1]int{23}

	fmt.Println("The original array is ", arr, " and it should be changed to ", [1]int{25}, " after mutation")

	mutateCopiedArray(arr)

	fmt.Println("The mutated array is (copied) ", arr)

	mutateOriginalArray(&arr)

	fmt.Println("The mutated array is ", arr)

}
