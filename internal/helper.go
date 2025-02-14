package internal

import (
	"fmt"
	"runtime"
)

func PrintErrorWithLine(msg string, err error) error {
	msg = msg + ": " + err.Error()

	pc, file, line, ok := runtime.Caller(1)
	if ok {
		fn := runtime.FuncForPC(pc)
		return fmt.Errorf("%s (file: %s, line: %d, func: %s)\n", msg, file, line, fn.Name())
	} else {
		return fmt.Errorf("%s\n", msg)
	}
}