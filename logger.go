package main

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// MAX_LOG_SIZE is the moximum size a file can grow to be before the
// log rotates to a new file.
const MAX_LOG_SIZE int64 = 80000

// LogLevel describes the level (type) of a log.
type LogLevel string

const (
	TraceLevel   LogLevel = "TRACE"
	InfoLevel    LogLevel = "INFO"
	WarningLevel LogLevel = "WARNING"
	ErrorLevel   LogLevel = "ERROR"
)

// Loggers global to entire applciation.
var (
	Trace          *log.Logger
	Info           *log.Logger
	Warning        *log.Logger
	Error          *log.Logger
	LogDestination [4]string
)

// InitLoggers uses the clinkn configuration object to initialize loggers.
func InitLoggers(c *ClinkConfig) (err error) {

	LogDestination = c.logDestination

	var traceHandler io.Writer
	var infoHandler io.Writer
	var warningHandler io.Writer
	var errorHandler io.Writer

	for i, v := range c.logDestination {
		var current io.Writer
		// var logf, logfn *os.File
		if v == "DISCARD" {
			current = ioutil.Discard
		} else if v == "STDERR" {
			current = os.Stderr
		} else if v == "STDOUT" {
			current = os.Stdout
		} else {
			_, fileName := determineCurrentLogFile(v)
			logfn, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return errors.New("logger: failed to create new log file")
			}
			current = logfn
		}

		if i == 0 {
			traceHandler = current
		} else if i == 1 {
			infoHandler = current
		} else if i == 2 {
			warningHandler = current
		} else if i == 3 {
			errorHandler = current
		}
	}
	if c.logLevel == "INFO" {
		traceHandler = ioutil.Discard
	} else if c.logLevel == "WARNING" {
		traceHandler = ioutil.Discard
		infoHandler = ioutil.Discard
	} else if c.logLevel == "ERROR" {
		traceHandler = ioutil.Discard
		infoHandler = ioutil.Discard
		warningHandler = ioutil.Discard
	}

	setLoggerHandles(traceHandler, infoHandler, warningHandler, errorHandler)

	return nil
}

// setLoggerHandles sets the writers and boilerplate data to display in each log message.
func setLoggerHandles(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	Trace = log.New(traceHandle, "TRACE: ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// determineCurrentLogFile reads the current directory to determine the current log file
// given a log name 'writerName'. If none exists returns false and the new file name to create.
func determineCurrentLogFile(writerName string) (bool, string) {
	ti := time.Now().UnixNano()
	t := strconv.Itoa(int(ti))

	files, _ := ioutil.ReadDir("./")
	var maxLtime int
	var maxFname string = ""
	for _, f := range files {
		if f.Size() < MAX_LOG_SIZE {
			fns := strings.Split(f.Name(), ".")
			if len(fns) == 3 {
				if fns[0] == writerName && fns[1] == "log" {
					if ltime, err := strconv.Atoi(fns[2]); err == nil {
						if ltime > maxLtime {
							maxLtime = ltime
							maxFname = f.Name()
						}
					}
				}
			}
		}
	}
	if maxFname == "" {
		return false, writerName + ".log." + t
	}
	return true, maxFname
}

// LogM logs a message to the appropriate level, checking and rolling log file if necessary.
func LogM(level LogLevel, message string) error {
	switch level {
	case TraceLevel:
		if LogDestination[0] != "DISCARD" && LogDestination[0] != "STDOUT" && LogDestination[0] != "STDERR" {
			//Then we are going to file
			_, fname := determineCurrentLogFile(LogDestination[0])
			logfile, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return errors.New("logger: failed to create or open new log file")
			}
			Trace.SetOutput(logfile)
		}
		Trace.Println(message)
	case InfoLevel:
		if LogDestination[1] != "DISCARD" && LogDestination[1] != "STDOUT" && LogDestination[1] != "STDERR" {
			//Then we are going to file
			_, fname := determineCurrentLogFile(LogDestination[1])
			logfile, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return errors.New("logger: failed to create or open new log file")
			}
			Info.SetOutput(logfile)
		}
		Info.Println(message)
	case WarningLevel:
		if LogDestination[2] != "DISCARD" && LogDestination[2] != "STDOUT" && LogDestination[2] != "STDERR" {
			//Then we are going to file
			_, fname := determineCurrentLogFile(LogDestination[2])
			logfile, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return errors.New("logger: failed to create or open new log file")
			}
			Warning.SetOutput(logfile)
		}
		Warning.Println(message)
	case ErrorLevel:
		if LogDestination[3] != "DISCARD" && LogDestination[3] != "STDOUT" && LogDestination[3] != "STDERR" {
			//Then we are going to file
			_, fname := determineCurrentLogFile(LogDestination[3])
			logfile, err := os.OpenFile(fname, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				return errors.New("logger: failed to create or open new log file")
			}
			Error.SetOutput(logfile)
		}
		Error.Println(message)
	}

	return nil
}
