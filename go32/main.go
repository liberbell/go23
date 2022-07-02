package main

import "fmt"

type UserNotFount struct {
	Username string
}

func (e *UserNotFount) Error() string {
	return fmt.Sprintf("User not fouond: %v", e.Username)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFount{Username: "mike"}
}

func main() {
	e1 := UserNotFount{"mike"}
	e2 := UserNotFount{"mike"}
	fmt.Println(e1 == e2)
	if err := myFunc(); err != nil {
		fmt.Println(err)
	}
}
