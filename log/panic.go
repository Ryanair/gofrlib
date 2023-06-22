package log

import "os"

func handlePanic() {
	if r := recover(); r != nil {
		log.Error("Panic occurred: %v", r)
		os.Exit(1)
	}
}
