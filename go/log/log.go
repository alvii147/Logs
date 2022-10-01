package log

import (
	"fmt"
	"strings"
	"time"
)

func Log(level LogLevel, v ...interface{}) {
	if !ENABLE_LOGGING {
		return
	}

	logElements := []interface{}{}

	timeFormat := strings.TrimSpace(fmt.Sprintf("[%s]", time.Now().Format(TIME_FORMAT)))

	prefixFormat := strings.TrimSpace(level.Prefix)
	if prefixFormat != "" {
		logElements = append(logElements, prefixFormat)
	}

	fileName, lineNum, funcName, _ := GetTrace()
	traceFormat := strings.TrimSpace(FormatTrace(fileName, lineNum, funcName))
	if traceFormat != "" {
		logElements = append(logElements, traceFormat)
	}

	messageFormat := strings.TrimSpace(fmt.Sprint(v...))
	if messageFormat != "" {
		logElements = append(logElements, messageFormat)
	}

	if !level.Suppressed {
		fmt.Fprintln(level.File, timeFormat, level.Colour.WrapText(logElements...))
	}
}

func LogError(v ...interface{}) {
	Log(LogLevelError, v...)
}

func LogWarning(v ...interface{}) {
	Log(LogLevelWarning, v...)
}

func LogSuccess(v ...interface{}) {
	Log(LogLevelSuccess, v...)
}

func LogDebug(v ...interface{}) {
	Log(LogLevelDebug, v...)
}

func LogInfo(v ...interface{}) {
	Log(LogLevelInfo, v...)
}

func LogNote(v ...interface{}) {
	Log(LogLevelNote, v...)
}
