package main

import (
	"errors"
	"fmt"
	"strings"
)

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
	// e1 := &UserNotFount{"mike"}
	// e2 := &UserNotFount{"mike"}
	// fmt.Println(e1 == e2)
	// if err := myFunc(); err != nil {
	// 	fmt.Println(err)
	// }

	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	e := errors.New("EOF")
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		// if err == io.EOF {
		if err == e {
			break
		}
	}
}
