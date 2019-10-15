package logs

import (
	"fmt"
	"time"
)

func Info(format string, a ...interface{}) {
	log("Info", format, a...)
}

func Error(format string, a ...interface{}) {
	log("Error", format, a...)
}

func log(level string, format string, a ...interface{}) {
	format = fmt.Sprintf("%s %s ", level, time.Now().Format("2006-01-02 15:04:05.000")) + format + "\n"
	fmt.Printf(format, a...)
}
