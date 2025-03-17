package errorutil

import (
	"fmt"
	"os"
)

type ErrorHandler struct {
	HadError bool
}

var GlobalErrorHandler = &ErrorHandler{HadError: false}

//...interface{} is likely accept one or more types of arguments
func ErrorUtil(line int, format string, args ...interface{}) {
	//args... is used to pass a slice to a variadic function
	message := fmt.Sprintf(format, args...)
	report(line, "", message)
	GlobalErrorHandler.HadError = true
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

