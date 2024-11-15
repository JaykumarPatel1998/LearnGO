package main

import "fmt"

type authorization struct {
	username string
	password string
}

func (auth authorization) authorize() string {
	return fmt.Sprintf("Authorization Basic -> %s : %s", auth.username, auth.password)
}

func main() {
	jay := authorization{
		username: "shaggy",
		password: "43jh5j4h4",
	}

	returnValue := jay.authorize()

	fmt.Println("The return value is", returnValue)
}
