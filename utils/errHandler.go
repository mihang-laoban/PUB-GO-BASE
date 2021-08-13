package utils

import (
	"log"
	"runtime"
)

func ErrorHandling(err error) {
	if err != nil {
		log.Printf("stack trace: %+v", err)
	}
}

func GetCallerInfo() (string, int) {
	_, file, line, _ := runtime.Caller(1)
	return file, line
}