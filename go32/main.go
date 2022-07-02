package main

import "fmt"

type UserNotFount struct {
	Username string
}

func (e *UserNotFount) Error() string {
	return e.Username
}

func myFunc() error {
	if ok {
		return nil
	}
	return &UserNotFount{Username: "mike"}
}

func main() {
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
