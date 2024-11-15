package main

import (
	"errors"
	"fmt"
	"strconv"
)

type user struct {
	name string
}

func getUser() (user, error) {
	user1 := user{
		name: "Jay",
	}
	if user1.name == "" {
		return user{}, errors.New("user not found")
	} else {
		return user1, nil
	}
}

func main() {
	user1, err := getUser()
	if err != nil {
		fmt.Println("Some Error Occured : ", err)
		return
	}
	fmt.Println("user value is ", user1)

	i, err := strconv.Atoi("432")
	if err != nil {
		fmt.Println("something went wront while converting ", err)
		return
	}

	fmt.Println("type casted string is ", i)
}
