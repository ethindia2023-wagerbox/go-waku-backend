package main

import (
	"fmt"
	"log"
)

func CheckErrorFatal(msg string, err error) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}

func CheckError(msg string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err)
		return
	}
}
