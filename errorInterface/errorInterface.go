package main

import "fmt"

type userError struct {
	name string
}

func (uError userError) Error() string {
	return fmt.Sprintf("Error thrown for user : %v", uError.name)
}

func sendMessage(userName string) error {
	can_send_message := true

	if !can_send_message {
		return userError{name: userName}
	} else {
		return nil
	}
}

func main() {
	err := sendMessage("Jay")
	if err != nil {
		fmt.Println("error : ", err)
	}

	fmt.Println("message sent")
}
