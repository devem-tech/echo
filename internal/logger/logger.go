package logger

import (
	"fmt"
	"os"
	"time"
)

type Logger struct {
	color     color
	isVerbose bool
}

func New(
	color color,
	isVerbose bool,
) *Logger {
	return &Logger{
		color:     color,
		isVerbose: isVerbose,
	}
}

func (l *Logger) Debug(format string, v ...any) {
	if !l.isVerbose {
		return
	}

	_, _ = fmt.Fprintln(os.Stdout, l.color.LightGray(l.ts()+" DEBUG : "+fmt.Sprintf(format, v...)))
}

func (l *Logger) Info(format string, v ...any) {
	_, _ = fmt.Fprintln(os.Stdout, l.color.LightGray(l.ts()+" INFO  : ")+fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(err error) {
	_, _ = fmt.Fprintln(os.Stderr, l.color.Red(l.ts()+" FATAL : "+err.Error()))

	os.Exit(1)
}

func (l *Logger) ts() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
