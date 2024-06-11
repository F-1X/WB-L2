package main

import (
	"errors"
	"fmt"
)

// Что выведет программа? Объяснить вывод программы.

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}
func test() *customError {
	{
		// do something
	}
	return nil
}
func main() {
	var err error
	err = test()
	if err != nil {
		// fmt.Printf("%+v", reflect.TypeOf(err))
		if errors.Is(err, &customError{}) {
			println("customError")
		}
		e := &customError{"ye"}
		fmt.Printf("%+v\n", e)

		if errors.As(err, *e) {
			println("customError")
		}

		if err.Error() == "customError" {
			println("customError")
		}
		println("error")
		return
	}
	println("ok")
}
