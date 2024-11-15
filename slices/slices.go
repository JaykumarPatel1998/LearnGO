package main

import "fmt"

func getMessages(messages []string) []float64 {
	length := len(messages)

	cost_slice := make([]float64, length)

	for i := 0; i < length; i++ {
		cost_slice[i] = 0.01 * float64(len(messages[i]))
	}
	return cost_slice
}

func main() {
	messages := []string{
		"hey this is a message",
		"but isn't this as well a message?",
		"indeed it is",
		"you are moody!",
	}

	costs := getMessages(messages)

	fmt.Println(costs)
}
