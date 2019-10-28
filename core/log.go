package core

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"sync"
	"time"
)

var mutex sync.Mutex
var projectRegex = regexp.MustCompile(`upmon/(.*)`)
var globalSyslogLevel = 3

// Log logs a message with the specified level
func log(level string, format string, values ...interface{}) {
	syslogLevel := getSyslogLogLevel(level)
	if syslogLevel > globalSyslogLevel {
		return
	}

	now := time.Now()
	formattedTime := now.Format("2006-01-02 15:04:05 MST")

	var color = 31
	if level == "debug" {
		color = 34
	} else if level == "notice" || level == "info" {
		color = 32
	}

	fpcs := make([]uintptr, 1)
	runtime.Callers(3, fpcs)

	caller := runtime.FuncForPC(fpcs[0] - 1)
	callerName := caller.Name()
	callerFile, callerLine := caller.FileLine(fpcs[0] - 1)
	callerFile = projectRegex.FindStringSubmatch(callerFile)[1]

	mutex.Lock()
	fmt.Fprintf(os.Stderr, "[\x1b[%dm%s\x1b[0m]\x1b[90m[%s]\x1b[90m[%d@%s][%s]\n", color, level, formattedTime, callerLine, callerFile, callerName)
	fmt.Fprintf(os.Stderr, "    └──\x1b[0m")
	fmt.Fprintf(os.Stderr, format, values...)
	fmt.Fprintf(os.Stderr, "\n")
	mutex.Unlock()
}

func getSyslogLogLevel(level string) int {
	switch level {
	case "emergency":
		return 0
	case "alert":
		return 1
	case "critical":
		return 2
	case "error":
		return 3
	case "warning":
		return 4
	case "notice":
		return 5
	case "info":
		return 6
	case "debug":
		return 7
	default:
		return 3
	}
}

// SetLogLevel sets the global log level
func SetLogLevel(level string) {
	globalSyslogLevel = getSyslogLogLevel(level)
}

// LogEmergency logs an emergency message
func LogEmergency(format string, values ...interface{}) {
	log("emergency", format, values...)
}

// LogAlert logs an alert message
func LogAlert(format string, values ...interface{}) {
	log("alert", format, values...)
}

// LogCritical logs a critical message
func LogCritical(format string, values ...interface{}) {
	log("critical", format, values...)
}

// LogError logs an error message
func LogError(format string, values ...interface{}) {
	log("error", format, values...)
}

// LogWarning logs a warning message
func LogWarning(format string, values ...interface{}) {
	log("warning", format, values...)
}

// LogNotice logs a notice message
func LogNotice(format string, values ...interface{}) {
	log("notice", format, values...)
}

// LogInfo logs an info message
func LogInfo(format string, values ...interface{}) {
	log("info", format, values...)
}

// LogDebug logs a debug message
func LogDebug(format string, values ...interface{}) {
	log("debug", format, values...)
}
