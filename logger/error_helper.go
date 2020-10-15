package logger

import (
	"fmt"
)

// ErrorHelper wraps the error logger
func ErrorHelper(err error) {
	Error(fmt.Sprintf("%T\n %s\n %#v\n", err, err, err))
}
