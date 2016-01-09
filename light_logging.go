package lightlogging

import (
	"log"
	"os"
	"runtime"
	"fmt"
)

func Info(message string) {
	green := "\033[0;32m"
	no_color := "\033[0m"

	log_message := fmt.Sprint("%s%s%s", green, message, no_color)

	var my_log = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	_, file, line, _ := runtime.Caller(1)

	my_log.Print("INFO:", file, ":", line, "-", log_message)
}

func Warning(message string) {
	yellow := "\033[0;33m"
	no_color := "\033[0m"
	log.Println(yellow + "WARNING: " + message + no_color)
}

func Error(message string) {
	red := "\033[0;31m"
	no_color := "\033[0m"
	log.Println(red + "ERROR: " + message + no_color)
}

func Critical(message string) {
	bold := "\033[1m"
	no_color := "\033[0m"
	log.Println(bold + "CRITICAL: " + message + no_color)
}
