package utils

import "log"

func ErrorHandling(err error) {
	if err != nil {
		log.Printf("stack trace: %+v", err)
	}
}