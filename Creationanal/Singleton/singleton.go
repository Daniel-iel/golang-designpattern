package main

import "fmt"

type MyLogger struct {
	loglevel int
}

func (l *MyLogger) Log(s string) {
	fmt.Println(l.loglevel, ":", s)
}

func (l *MyLogger) SetlogLevel(level int) {
	l.loglevel = level
}

var logger *MyLogger

func getLoggerInstance() *MyLogger {
	if logger == nil {
		fmt.Println("Creating Logger Instance")
		logger = &MyLogger{}
	}
	fmt.Println("Returning Logger Instance")
	return logger
}
