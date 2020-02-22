package log

import (
	"fmt"
	slog "log"
	"os"
	"runtime"
	"strings"
)

const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)

var level = InfoLevel
var flags = LstdFlags | Lshortfile

func SetLevel(lvl Level) {
	level = lvl
}

func SetFlags(flag int) {
	if (flag & Lshortfile) != 0 {
		flag -= Lshortfile
	} else if (flag & Llongfile) != 0 {
		flag -= Llongfile
	}
	flags = flag
	slog.SetFlags(flag)
}

func Debugf(format string, args ...interface{}) {
	logf(DebugLevel, format, args...)
}

func Infof(format string, args ...interface{}) {
	logf(InfoLevel, format, args...)
}

func Printf(format string, args ...interface{}) {
	logf(InfoLevel, format, args...)
}

func Warnf(format string, args ...interface{}) {
	logf(WarnLevel, format, args...)
}

func Errorf(format string, args ...interface{}) {
	logf(ErrorLevel, format, args...)
}

func Fatalf(format string, args ...interface{}) {
	logf(FatalLevel, format, args...)
	os.Exit(1)
}

func Panicf(format string, args ...interface{}) {
	logf(PanicLevel, format, args...)
	os.Exit(1)
}

func Debug(args ...interface{}) {
	log(DebugLevel, args...)
}

func Info(args ...interface{}) {
	log(InfoLevel, args...)
}

func Print(args ...interface{}) {
	log(InfoLevel, args...)
}

func Warn(args ...interface{}) {
	log(WarnLevel, args...)
}

func Error(args ...interface{}) {
	log(ErrorLevel, args...)
}

func Fatal(args ...interface{}) {
	log(FatalLevel, args...)
	os.Exit(1)
}

func Panic(args ...interface{}) {
	log(PanicLevel, args...)
	os.Exit(1)
}

func Println(args ...interface{}) {
	log(InfoLevel, args...)
}

func logf(lvl Level, format string, args ...interface{}) {
	if lvl > level {
		return
	}

	if (flags & Lshortfile) != 0 || (flags & Llongfile) != 0 {
		_, file, line, _ := runtime.Caller(2)
		if (flags & Lshortfile) != 0 {
			pos := strings.LastIndex(file, "/")
			if pos >= 0 {
				file = file[pos+1:]
			}
		}
		format = fmt.Sprintf("%s:%d: %v %s", file, line, lvl, format)
	} else {
		format = fmt.Sprintf("%v %s", lvl, format)
	}

	slog.Printf(format, args...)
}

func log(lvl Level, args ...interface{}) {
	if lvl > level {
		return
	}

	var largs []interface{}
	if (flags & Lshortfile) != 0 || (flags & Llongfile) != 0 {
		_, file, line, _ := runtime.Caller(2)
		if (flags & Lshortfile) != 0 {
			pos := strings.LastIndex(file, "/")
			if pos >= 0 {
				file = file[pos+1:]
			}
		}
		largs = append(largs, fmt.Sprintf("%s:%d: %v", file, line, lvl))
	} else {
		largs = append(largs, lvl)
	}
	largs = append(largs, args...)
	slog.Println(largs...)
}