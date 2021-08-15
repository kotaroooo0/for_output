package main

import (
	"fmt"
)

func GetSomething() (int, error) {
	fmt.Println("GetSomething()")
	// return 0, &AppError{
	// 	Code:    "NOTFOUND",
	// 	Message: "hoge",
	// }
	return 0, Wrap(nil)
}

func main() {
	fmt.Println("Start")
	_, err := GetSomething()
	if aerr, ok := err.(*AppError); ok {
		switch aerr.Code {
		case "NOTFOUND":
			fmt.Println("error code is NOTFOUND")
		default:
			fmt.Println("error is not nil")
		}
	}
	fmt.Println("End")
}

type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	fmt.Println("Error()")
	return fmt.Sprintf("code: %s message: %s", e.Code, e.Message)
}

func Wrap(err error) error {
	fmt.Println("Wrap(err)")
	if err == nil {
		return nil
	}
	return &AppError{
		Message: err.Error(),
	}
}
