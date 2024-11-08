package logger

import (
	"log"
	"os"
)

func PrintInfo(msg string) {
	log.Printf("INFO: %s\n", msg)
}

func PrintError(err error) {
	log.Printf("ERROR: %s\n", err.Error())
}

func FatalError(err error) {
	PrintError(err)
	os.Exit(1)
}
