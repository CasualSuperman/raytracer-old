package log

import "log"
import "os"

var logger = log.New(os.Stderr, "", 0)

/* Fatal functions call Print, then exit(1) */
func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatalf(format, v...)
}

func Fatalln(v ...interface{}) {
	logger.Fatalln(v...)
}

/* Panic functions call print, then panic() */
func Panic(v ...interface{}) {
	logger.Panic(v...)
}

func Panicf(format string, v ...interface{}) {
	logger.Panicf(format, v...)
}

func Panicln(v ...interface{}) {
	logger.Panicln(v...)
}

/* Print functions print. */
func Print(v ...interface{}) {
	logger.Print(v...)
}

func Printf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

func Println(v ...interface{}) {
	logger.Println(v...)
}
