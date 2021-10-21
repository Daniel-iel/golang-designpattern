package main

import (
	"fmt"
)

func main() {
	log := getLoggerInstance()
	log.SetlogLevel(1)
	log.Log("This is a log message")

	log := getLoggerInstance()
	log.SetlogLevel(2)
	log.Log("This is a log message")

	log := getLoggerInstance()
	log.SetlogLevel(3)
	log.Log("This is a log message")

	for i := 1; i < 10; i++ {
		go getLoggerInstance()
	}
	fmt.Scanln()
}
