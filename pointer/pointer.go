package main

import "fmt"

func increment(number *int) {
	fmt.Println("address of number is : ", number)
	*number = *number + 1
	fmt.Println("value of dereferenced number is : ", *number)
}

func main() {
	number := 23
	pass := &number
	increment(pass)
	fmt.Printf("incremented value of number is %d", number)
}
