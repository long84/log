package log

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	Ldatetime	= 1 << iota
	Lfile
)

type Logger struct {
	sync.Mutex
	flag		int
	level		Level
	prefix		string
	out			io.Writer
	buf   		[]byte
	callDepth 	int
}

func (l *Logger) SetFlags(flag int) {
	l.Lock()
	defer l.Unlock()
	l.flag = flag
}

func (l *Logger) SetLevel(level Level) {
	l.Lock()
	defer l.Unlock()
	l.level = level
}

func (l *Logger) SetPrefix(prefix string) {
	l.Lock()
	defer l.Unlock()
	l.prefix = prefix
}

func (l *Logger) SetOutput(out io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.out = out
}

func (l *Logger) SetCallDepth(depth int) {
	l.Lock()
	defer l.Unlock()
	l.callDepth = depth
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logf(DebugLevel, format, args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.logf(InfoLevel, format, args...)
}

func (l *Logger) Printf(format string, args ...interface{}) {
	l.logf(InfoLevel, format, args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logf(WarnLevel, format, args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logf(ErrorLevel, format, args...)
}

func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logf(FatalLevel, format, args...)
	os.Exit(1)
}

func (l *Logger) Panicf(format string, args ...interface{}) {
	l.logf(PanicLevel, format, args...)
	os.Exit(1)
}

func (l *Logger) Debug(args ...interface{}) {
	l.log(DebugLevel, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.log(InfoLevel, args...)
}

func (l *Logger) Print(args ...interface{}) {
	l.log(InfoLevel, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.log(WarnLevel, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.log(ErrorLevel, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.log(FatalLevel, args...)
	os.Exit(1)
}

func (l *Logger) Panic(args ...interface{}) {
	l.log(PanicLevel, args...)
	os.Exit(1)
}

func (l *Logger) Println(args ...interface{}) {
	l.log(InfoLevel, args...)
}

func (l *Logger) logf(level Level, format string, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if level > l.level {
		return
	}

	l.output(level, fmt.Sprintf(format, args...))
}

func (l *Logger) log(level Level, args ...interface{}) {
	l.Lock()
	defer l.Unlock()

	if level > l.level {
		return
	}

	l.output(level, fmt.Sprintln(args...))
}

func (l *Logger) output(level Level, s string) {
	l.buf = l.buf[:0]
	l.formatHeader(&l.buf, level)
	l.buf = append(l.buf, s...)
	if len(s) == 0 || s[len(s)-1] != '\n' {
		l.buf = append(l.buf, '\n')
	}
	l.out.Write(l.buf)
}

func (l *Logger) formatHeader(buf *[]byte, level Level) {
	if (l.flag & Ldatetime) != 0 {
		*buf = append(*buf, []byte(time.Now().Format("2006-01-02 15:04:05"))...)
		*buf = append(*buf, ' ')
	}

	*buf = append(*buf, []byte(level.String())...)
	*buf = append(*buf, ' ')

	if (l.flag & Lfile) != 0 {
		_, file, line, ok := runtime.Caller(l.callDepth+3)
		if !ok {
			file = "???"
		}
		//文件名保留两级
		first := true
		for i := len(file) - 1; i > 0; i-- {
			if file[i] == '/' {
				if first {
					first = false
					continue
				}
				file = file[i+1:]
				break
			}
		}
		*buf = append(*buf, file...)
		*buf = append(*buf, ':')
		*buf = append(*buf, fmt.Sprintf("%d: ", line)...)
	}

	if l.prefix != "" {
		*buf = append(*buf, []byte(l.prefix)...)
	}
}

func New(out io.Writer, prefix string, flag int) *Logger {
	return &Logger{out: out, prefix: prefix, flag: flag, level: InfoLevel, callDepth: 1}
}