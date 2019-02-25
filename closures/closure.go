package closures

import (
	"fmt"
	"time"
)

func LogLater(message string) func() string {
	//Variables declared either in a function body, or as a function argument are available in inline function declarations
	timestamp := time.Now()
	return func() string {
		return fmt.Sprintf("%s from %s", message, timestamp.String())
	}
}

func AnnotatedLogger(annotation string) func(string) {
	//Sometimes we only want to store some variables before hand, and allow others to be passed in later
	return func(message string) {
		fmt.Printf("%s: %s", annotation, message)
	}
}
