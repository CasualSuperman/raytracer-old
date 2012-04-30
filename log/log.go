// Implements a logging package attached to Go's builtin logging package, and
// all input goes to stderr.
package log

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 0)

// Prints the given arguments to stderr, followed by a call to
// os.Exit(1)
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

// Prints the given arguments to stderr using the provided format specifier,
// followed by a call to os.Exit(1)
func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

// Prints the given arguments to stderr followed by a newline, then calls
// os.Exit(1)
func Fatalln(v ...interface{}) {
	logger.Fatalln(v...)
}

// Prints the given arguments to stderr.
func Print(v ...interface{}) {
	logger.Print(v...)
}

// Prints the given arguments to stderr, using the provided format specifier.
func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

// Prints the given arguments to stderr, followed by a newline.
func Println(v ...interface{}) {
	logger.Println(v...)
}
