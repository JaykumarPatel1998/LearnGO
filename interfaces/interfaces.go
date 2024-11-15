package main

import (
	"fmt"
	"time"
)

// A type can implement multiple interfaces
type message interface {
	getMessage() string
}

type Copier interface {
	Copy(sourceFile string, destinationFile string) int
}

type birthdayMessage struct {
	dob       time.Time
	recipient string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("happy birthday : %s", bm.recipient)
}

func sendMessage(msg message) {
	fmt.Println(msg.getMessage())
}

func main() {
	bm := birthdayMessage{
		dob:       time.Now(),
		recipient: "Jay",
	}
	sendMessage(bm)
}
