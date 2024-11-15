package main

import (
	"fmt"
)

type messageToSend struct {
	sender    user
	recipient user
	message   string
}

type user struct {
	name   string
	number int
}

func canSendMessage(msg messageToSend) bool {
	if msg.sender.name == "" {
		return false
	}
	if msg.recipient.name == "" {
		return false
	}
	if msg.sender.number == 0 {
		return false
	}
	if msg.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	msg_sender := user{
		name:   "Jay",
		number: 322111,
	}
	msg_sender.name = "Jaykumar Patel"

	msg_recipient := user{
		name:   "Vilas",
		number: 9273894,
	}

	msg := messageToSend{
		sender:    msg_sender,
		recipient: msg_recipient,
		message:   "Hey what's up?",
	}

	if returnValue := canSendMessage(msg); returnValue {
		fmt.Printf("message sent : %v", msg)
	} else {
		fmt.Println("message not sent...")
	}
}
