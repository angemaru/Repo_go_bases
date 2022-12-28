package main

import (
	"fmt"
)

type customError struct {
	msg  string
	code int
}

func (c customError) Error() string {
	return c.msg + fmt.Sprintf(" msg: %s, code: %d", c.msg, c.code)
}

func _main() {
	err := &customError{
		msg:  "something has failed",
		code: 500,
	}
	fmt.Println("error", err)
}

/*
type customError struct {
	message string
}

func (c *customError) Error() string {
	return c.message
}

func NewCustom(message string) error {
	return &customError{message: message}
}

var (
	ErrNotFound = errors.New("error: item was not found")
	ErrUnknown  = errors.New("general error")
	ErrOther    = NewCustom("error item was not found")
)

func main() {
	err := SendError(true)

	errdetail := fmt.Errorf("%w, %s", err, "here are details")

	fmt.Println("error", errdetail)
	fmt.Println(errors.Is(errdetail, ErrNotFound))

}

func SendError(b bool) error {
	if b {
		return ErrNotFound

	}
	return nil
}
*/
/*
Is()
type customError struct {
	message string
}

func (c *customError) Error() string{
	return c.message
}

func NewCustom(message string) error{
	return &customError{message: message}
}
var (
	ErrNotFound = errors.New("error: item was not found")
	ErrUnknown  = errors.New("general error")
	ErrOther = NewCUstom("error item was not found")
)

func main() {
	err := SendError(true)

	fmt.Printf(error.Is(err, ErrNotFound))
	fmt.Printf(error.Is(err, ErrOther))
	fmt.Printf(err == ErrNotFound)
}

func SendError(b bool) error {
	if b {
		return ErrNotFound

	}
*/

/*
var (
	ErrNotFound = errors.New("error: item was not found")
	ErrUnknown  = errors.New("general error")
)

func main() {
	err := SendError(true)

	fmt.Printf("%T", err)
}

func SendError(b bool) error {
	if b {
		return ErrNotFound

	}
	return nil
}
*/
