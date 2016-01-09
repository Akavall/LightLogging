package light_logging

import (
	"log"
	"os"
	"runtime"
	"fmt"
)

func Info(message string) {
	green := "\033[0;32m"
	no_color := "\033[0m"

	var my_log = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	_, file, line, _ := runtime.Caller(1)

	log_message := fmt.Sprint(green, "INFO:", file, ":", line, "-", message, no_color)

	my_log.Print(log_message)
}


func Warning(message string) {
	yellow := "\033[0;33m"
	no_color := "\033[0m"

	var my_log = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	_, file, line, _ := runtime.Caller(1)

	log_message := fmt.Sprint(yellow, "WARNING:", file, ":", line, "-", message, no_color)

	my_log.Print(log_message)
}

func Error(message string) {
	red := "\033[0;31m"
	no_color := "\033[0m"

	var my_log = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	_, file, line, _ := runtime.Caller(1)

	log_message := fmt.Sprint(red, "ERROR:", file, ":", line, "-", message, no_color)

	my_log.Print(log_message)
}

func Critical(message string) {
	bold := "\033[1m"
	no_color := "\033[0m"

	var my_log = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	_, file, line, _ := runtime.Caller(1)

	log_message := fmt.Sprint(bold, "CRITICAL:", file, ":", line, "-", message, no_color)

	my_log.Fatal(log_message)
}
