package log

import "io/ioutil"
import "flag"
import "log"
import "os"

var Logger *log.Logger

var doLog = flag.Bool("debug", false, "Turn on debug information.")

func init() {
	flag.Parse()

	if *doLog {
		Logger = log.New(os.Stderr, "", log.Lmicroseconds | log.Lshortfile)
	} else {
		Logger = log.New(ioutil.Discard, "", 0)
	}
}
