package util

import (
	"log"
)

func FatalIfError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func FatalIfTrue(condition bool, message string) {
	if condition {
		log.Fatal(message)
	}
}
