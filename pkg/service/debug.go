package service

import (
	"fmt"
	"log"
)

func debug(format string, v ...interface{}) {
	if true {
		format = fmt.Sprintf("[debug] %s\n", format)
		_ = log.Output(2, fmt.Sprintf(format, v...))
	}
}
