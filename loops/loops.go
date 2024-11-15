package main

import "fmt"

// func fizzBuzz(n int) []string {
// 	var response []string

// 	for i := 1; i <= n; i++ {
// 		if i%3 == 0 {
// 			response = append(response, "Fizz")
// 		} else if i%5 == 0 {
// 			response = append(response, "Buzz")
// 		} else if i%3 == 0 && i%5 == 0 {
// 			response = append(response, "FizzBuzz")
// 		} else {
// 			response = append(response, string(i))
// 		}
// 	}

// 	return response
// }

func bulkSend(numMessages int) float64 {
	totalCost := 0.0

	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func main() {
	result := bulkSend(4)

	if result != 4.06 {
		fmt.Println("Invalid operation!!")
		return
	} else {
		fmt.Println("Valid Operation")
	}
}
