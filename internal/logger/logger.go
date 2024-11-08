package logger

import "log"

func PrintInfo(msg string) {
	log.Printf("INFO: %s\n", msg)
}

func PrintError(err error) {
	log.Printf("ERROR: %s\n", err.Error())
}
