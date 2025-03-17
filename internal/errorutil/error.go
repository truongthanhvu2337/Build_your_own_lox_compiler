package errorutil

import (
	"fmt"
	"os"
)

type ErrorHandler struct {
	HadError bool
}

var GlobalErrorHandler = &ErrorHandler{HadError: false}

func ErrorUtil(line int, format string, args ...interface{}) {
	message := fmt.Sprintf(format, args...)
	report(line, "", message)
	GlobalErrorHandler.HadError = true
}

func report(line int, where, message string) {
	fmt.Fprintf(os.Stderr, "[line %d] Error%s: %s\n", line, where, message)
}

