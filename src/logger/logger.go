package logger

import (
	"fmt"
	"os"
	"strings"
)

var debug = strings.ToLower(os.Getenv("DEBUG"))

func Debug(message string) {
	if debug == "true" {
		fmt.Println("[DEBUG] " + message)
	}
}

func Info(message string) {
	fmt.Println("[INFO] " + message)
}
