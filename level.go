package log

import (
	"fmt"
	"strings"
)

type Level uint32

const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

func (level Level) String() string {
	switch level {
	case PanicLevel:
		return "[PANIC]"
	case FatalLevel:
		return "[FATAL]"
	case ErrorLevel:
		return "[ERROR]"
	case WarnLevel:
		return "[WARN]"
	case InfoLevel:
		return "[INFO]"
	case DebugLevel:
		return "[DEBUG]"
	default:
		return "[UNKNOWN]"
	}
}

// ParseLevel takes a string level and returns the Logrus log level constant.
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid log Level: %q", lvl)
}

