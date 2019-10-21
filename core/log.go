package core

import (
  "sync"
  "fmt"
  "time"
  "runtime"
  "regexp"
)

var mutex sync.Mutex
var projectRegex = regexp.MustCompile(`upmon/(.*)`)

// Log logs a message with the specified level
func log(level string, format string, values ...interface{}) {
  now := time.Now()
  formattedTime := now.Format("2006-01-02 15:04:05 MST")

  var color = 31
  if level == "debug" {
    color = 32
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
  fmt.Printf("[\x1b[%dm%s\x1b[0m]\x1b[90m[%s]\x1b[90m[%d@%s][%s]\n", color, level, formattedTime, callerLine, callerFile, callerName)
  fmt.Printf("    └──\x1b[0m")
  fmt.Printf(format, values...)
  fmt.Println()
  mutex.Unlock()
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
